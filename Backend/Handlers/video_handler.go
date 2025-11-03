package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO implement video processing (e.g., change metadata so it's suitable for streaming, thumbnail generation)
func handleVideoProcessing(videoPath string) {
	println("Processing video:", videoPath)
}

func UploadVideoHandler(c *gin.Context) {
	println("something")
	file, err := c.FormFile("videoFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - User ID not found"})
		return
	}

	userIDStr := fmt.Sprintf("%d", userID)
	videoDir := "./Uploads/TempVideoPath/"

	filename := fmt.Sprintf("%s%s___%s", videoDir, userIDStr, file.Filename)
	println("fileName: ", filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	go handleVideoProcessing(filename)
	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully"})
}
