package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	utils "VideoStreamingBackend/Utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func LoginHandler(c *gin.Context) {
	var input LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := utils.CreateToken(user.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create token"})
		return
	}
	refreshTokenString, expiresAt, err := utils.CreateRefreshToken(user.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create refresh token"})
		return
	}

	refreshToken := models.RefreshToken{
		UserID:    user.UserID,
		Token:     refreshTokenString,
		ExpiresAt: expiresAt,
	}

	if err := config.DB.Create(&refreshToken).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create refresh token"})
		return
	}

	c.SetCookie("token", tokenString, 3600*24, "/", "localhost", false, false)
	c.SetCookie("refresh_token", refreshTokenString, 3600*24*7, "/", "localhost", false, false)

	c.JSON(200, gin.H{"token": tokenString, "user": user.UserName, "refreshToken": refreshTokenString})
}

/* TODO: opdeter api doc hvis der er tid til det */
func LogoutHandler(c *gin.Context) {
	/* result := revokeRefreshToken(c.GetHeader("Authorization")[7:]) */
	type LogoutRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	var req LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := revokeRefreshToken(req.RefreshToken)
	if result != nil {
		c.JSON(500, gin.H{"error": "Failed to revoke refresh token", "details": result.Error()})
		return
	}
	/* c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	c.SetCookie("token", "", -1, "/", "localhost", false, true) */
	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

func revokeRefreshToken(tokenString string) error {
	now := time.Now()
	println("Token:", tokenString)

	result := config.DB.Model(&models.RefreshToken{}).
		Where("token = ?", tokenString).
		Update("revoked", false).
		Update("revoked_at", now)

	if result.Error != nil {
		println("Error revoking token:", result.Error.Error())
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no refresh token found to revoke")
	}
	return nil
}
