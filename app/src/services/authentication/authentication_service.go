package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/app/src/controllers/authentication"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/app/src/models"
	validator "github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/configs"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils/types"
	"strconv"
)

func AddUser(userDetails *authentication.SignUpDto) (*models.User, error) {
	user := models.User{
		Email:        userDetails.Email,
		PasswordHash: utils.GeneratePassword(userDetails.Password),
		UserStatus:   types.ActiveUser,
		UserRole:     userDetails.UserRole,
	}
	if err := validator.ValidateRequest(user); err != nil {
		return nil, err
	}

	result := configs.DbClient.Create(user)
	if (result.Error != nil) || (result.RowsAffected == utils.IntZero) {
		return nil, errors.New(fmt.Sprintf("error adding user: %s", result.Error))
	}

	user.PasswordHash = ""
	return &user, nil
}

func GetUserByEmail(ctx context.Context, userEmail string) (*models.User, error) {
	user := &models.User{}
	result := configs.DbClient.WithContext(ctx).Model(user).Where("email = ?", userEmail).First(user)
	return user, result.Error
}

func VerifyUser(ctx context.Context, userDetails *authentication.SignInDto) (*utils.Tokens, error) {
	user, err := GetUserByEmail(ctx, userDetails.Email)
	if err != nil {
		return nil, err
	}
	userId := strconv.Itoa(int(user.ID))

	isValidUser := utils.ComparePasswords(user.PasswordHash, userDetails.Password)
	if !isValidUser {
		return nil, errors.New("wrong user email address or password")
	}

	tokens, err := utils.GenerateNewTokens(userId, user.Email, user.UserRole, nil)
	if err != nil {
		return nil, err
	}

	err = configs.RedisClient.Set(ctx, userId, tokens.Refresh, 0).Err()
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
