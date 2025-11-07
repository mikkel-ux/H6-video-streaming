package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	DTO "VideoStreamingBackend/Models/DTO"
	utils "VideoStreamingBackend/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(g *gin.Context) {
	var req DTO.CreateUserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HasshPassword(req.Password)
	if err != nil {
		g.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
		Age:       req.Age,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		g.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	channel := models.Channel{
		Name:        req.ChannelName,
		Description: req.ChannelDescription,
		UserID:      user.UserID,
	}

	config.DB.Model(&user).Association("Channel").Append(&channel)

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "userId": user.UserID})
}

func GetUserHandler(g *gin.Context) {
	userId := g.Param("userId")
	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var channel models.Channel
	if err := config.DB.Where("user_id = ?", user.UserID).First(&channel).Error; err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Channel not found"})
		return
	}

	g.JSON(http.StatusOK, DTO.GetUserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		UserName:  user.UserName,
		Age:       user.Age,
		ChannelID: channel.ChannelID,
	})
}

func DeleteUserHandler(g *gin.Context) {
	userId := g.Param("userId")

	result := config.DB.Delete(&models.User{}, userId)
	if result.Error != nil {
		g.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	g.JSON(200, gin.H{"message": "User deleted successfully"})
}

func UpdateUserHandler(g *gin.Context) {
	userId := g.Param("userId")
	var req DTO.UpdateUserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		g.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.UserName != nil {
		user.UserName = *req.UserName
	}
	if req.Age != nil {
		user.Age = *req.Age
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	if err := config.DB.Save(&user).Error; err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(200, gin.H{"message": "User updated successfully"})
}

func UpdatePasswordHandler(g *gin.Context) {
	userId := g.Param("userId")
	type UpdatePasswordRequest struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required,min=8"`
	}
	var req UpdatePasswordRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		g.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if !utils.CheckPasswordHash(req.OldPassword, user.Password) {
		g.JSON(401, gin.H{"error": "Old password is incorrect"})
		return
	}

	hashedNewPassword, err := utils.HasshPassword(req.NewPassword)
	if err != nil {
		g.JSON(500, gin.H{"error": "Failed to hash new password"})
		return
	}
	user.Password = hashedNewPassword

	if err := config.DB.Save(&user).Error; err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(200, gin.H{"message": "Password updated successfully"})
}
