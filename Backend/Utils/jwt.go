package utils

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Minute * 30).Unix(),
		},
	)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func CreateRefreshToken(userID int64) (string, time.Time, error) {
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     expiresAt.Unix(),
		},
	)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userID := token.Claims.(jwt.MapClaims)["user_id"].(float64)

	tokenIsInDB := config.DB.Where("token = ? AND userID = ? AND revoked = ?", tokenString, userID, false).First(&models.RefreshToken{})
	if tokenIsInDB.Error != nil {
		return nil, fmt.Errorf("refresh token not found or revoked")
	}

	return token, nil
}
