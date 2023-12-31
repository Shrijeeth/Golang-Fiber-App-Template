package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/models"
	validator "github.com/Shrijeeth/Golang-Fiber-App-Template/pkg"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/jobs"
	"github.com/gocraft/work"
	"gorm.io/gorm"
	"strconv"
)

func AddUser(userDetails *types.AddUserData) (*models.User, error) {
	var user *models.User

	err := configs.DbClient.Transaction(func(tx *gorm.DB) error {
		if (userDetails.Password != "") && (userDetails.AuthenticationType == types.NormalAuthentication) {
			user = &models.User{
				Username:           userDetails.Username,
				Email:              userDetails.Email,
				PasswordHash:       utils.GeneratePassword(userDetails.Password),
				UserStatus:         types.ActiveUser,
				UserRole:           userDetails.UserRole,
				AuthenticationType: userDetails.AuthenticationType,
			}
		} else if (userDetails.Password == "") && (userDetails.AuthenticationType == types.GoogleAuthentication) {
			user = &models.User{
				Username:           userDetails.Username,
				Email:              userDetails.Email,
				UserStatus:         types.ActiveUser,
				UserRole:           types.User,
				AuthenticationType: userDetails.AuthenticationType,
			}
		}

		if err := validator.ValidateRequest(user); err != nil {
			return err
		}

		result := tx.Create(user)
		if (result.Error != nil) || (result.RowsAffected == utils.IntZero) {
			tx.Rollback()
			return fmt.Errorf("error adding user: %s", result.Error)
		}

		user.PasswordHash = ""

		_, err := jobs.JobEnqueuer.Enqueue(jobs.SampleMailJobName, work.Q{
			"username": user.Username,
			"email":    user.Email,
		})
		if err != nil {
			return fmt.Errorf("error executing email job: %s", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{}
	result := configs.DbClient.Model(user).Where("email = ?", userEmail).First(user)
	return user, result.Error
}

func VerifyUser(ctx context.Context, userDetails *types.VerifyUserData) (*utils.Tokens, error) {
	var tokens *utils.Tokens

	err := configs.DbClient.Transaction(func(tx *gorm.DB) error {
		user, err := GetUserByEmail(userDetails.Email)
		if err != nil {
			tx.Rollback()
			return err
		}
		userId := strconv.Itoa(int(user.ID))

		if user.AuthenticationType == types.NormalAuthentication {
			isValidUser := utils.ComparePasswords(user.PasswordHash, userDetails.Password)
			if !isValidUser {
				tx.Rollback()
				return errors.New("wrong user email address or password")
			}
		}

		tokens, err = utils.GenerateNewTokens(userId, user.Email, user.UserRole, nil)
		if err != nil {
			tx.Rollback()
			return err
		}

		if configs.IsCacheRequired() {
			err = configs.RedisClient.Set(ctx, userId, tokens.Refresh, 0).Err()
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tokens, nil
}
