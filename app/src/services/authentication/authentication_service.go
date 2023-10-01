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
	"gorm.io/gorm"
	"os"
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

		if configs.IsMailClientRequired() {
			from := types.MailDetails{
				Email: os.Getenv("SMTP_EMAIL_ID"),
				Name:  "Go-Fiber Template",
			}
			to := []types.MailDetails{
				{
					Name:  user.Username,
					Email: user.Email,
				},
			}
			templateVariables := map[string]string{
				"name": user.Username,
			}
			err := configs.MailClient.SendMailWithTemplate("Registration Success", "sample-template.html", from, to, templateVariables)
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
