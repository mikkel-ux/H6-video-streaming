package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	utils "VideoStreamingBackend/Utils"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	UserName  string `json:"userName" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

func CreateUserHandler(g *gin.Context) {
	var req CreateUserRequest
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

	g.JSON(201, gin.H{"message": "User created successfully", "userId": user.UserID})
}

type GetUserResponse struct {
	UserID   uint   `json:"userId"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Age      int    `json:"age"`
}

func GetUserHandler(g *gin.Context) {
	userId := g.Param("userId")
	var user models.User
	if err := config.DB.First(&GetUserResponse{}, userId).Error; err != nil {
		g.JSON(404, gin.H{"error": "User not found"})
		return
	}

	g.JSON(200, user)
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

type UpdateUserRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	UserName  *string `json:"userName,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func UpdateUserHandler(g *gin.Context) {
	userId := g.Param("userId")
	var req UpdateUserRequest
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
