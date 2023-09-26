package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"strings"
	"time"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id string, email string, userRole types.UserRole, additionalDetails map[string]interface{}) (*Tokens, error) {
	accessToken, err := generateNewAccessToken(id, email, userRole, additionalDetails)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string, email string, userRole types.UserRole, additionalDetails map[string]interface{}) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["email"] = email
	claims["user_role"] = userRole

	// Set additional credentials:
	for key, value := range additionalDetails {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
