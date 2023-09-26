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
	"strconv"
)

func AddUser(userDetails *SignUpDto) (*models.User, error) {
	var user *models.User

	err := configs.DbClient.Transaction(func(tx *gorm.DB) error {
		user = &models.User{
			Email:        userDetails.Email,
			PasswordHash: utils.GeneratePassword(userDetails.Password),
			UserStatus:   types.ActiveUser,
			UserRole:     userDetails.UserRole,
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
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(tx *gorm.DB, userEmail string) (*models.User, error) {
	user := &models.User{}
	result := tx.Model(user).Where("email = ?", userEmail).First(user)
	return user, result.Error
}

func VerifyUser(ctx context.Context, userDetails *SignInDto) (*utils.Tokens, error) {
	var tokens *utils.Tokens

	err := configs.DbClient.Transaction(func(tx *gorm.DB) error {
		user, err := GetUserByEmail(tx, userDetails.Email)
		if err != nil {
			tx.Rollback()
			return err
		}
		userId := strconv.Itoa(int(user.ID))

		isValidUser := utils.ComparePasswords(user.PasswordHash, userDetails.Password)
		if !isValidUser {
			tx.Rollback()
			return errors.New("wrong user email address or password")
		}

		tokens, err = utils.GenerateNewTokens(userId, user.Email, user.UserRole, nil)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = configs.RedisClient.Set(ctx, userId, tokens.Refresh, 0).Err()
		if err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tokens, nil
}
