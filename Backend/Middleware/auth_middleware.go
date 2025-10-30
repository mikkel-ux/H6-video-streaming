package middleware

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	utils "VideoStreamingBackend/Utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	result := config.DB.Model(&models.RefreshToken{}).
		Where("token = ? AND revoked = false AND expires_at > ?", refreshTokenString, time.Now()).
		First(&models.RefreshToken{})
	if result.Error != nil {
		return "", result.Error
	}
	newAccessToken, err := utils.CreateToken(claims.UserID)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			refreshToken, err := c.Cookie("refresh_token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authentication provided"})
				return
			}

			newAccessToken, err := RefreshAccessToken(refreshToken)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to refresh access token"})
				return
			}

			c.Header("Authorization", "Bearer "+newAccessToken)
			authHeader = "Bearer " + newAccessToken
		}
		var accessToken string
		if strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = authHeader[7:]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		println(accessToken) //TODO remove debug print

		c.Next()
	}
}
