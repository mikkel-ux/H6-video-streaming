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
