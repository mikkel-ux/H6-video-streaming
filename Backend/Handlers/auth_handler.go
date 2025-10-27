package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	utils "VideoStreamingBackend/Utils"

	"github.com/gin-gonic/gin"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func LoginUserHandler(g *gin.Context) {
	var req LoginUserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		g.JSON(401, gin.H{"error": "Invalid email or password 1"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		g.JSON(401, gin.H{"error": "Invalid email or password 2"})
		return
	}

	g.JSON(200, gin.H{"message": "Login successful", "userId": user.UserID})
}
