/* https://superuser.com/questions/1682333/ffmpeg-how-can-i-get-the-first-frame-of-an-mp4-and-maintain-its-aspect-ratio img from video
// using ffmpeg */

package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	"fmt"
	"net/http"

	DTO "VideoStreamingBackend/Models/DTO"
	utils "VideoStreamingBackend/Utils"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
)

func UploadVideoHandler(c *gin.Context) {
	var request DTO.UploadVideoRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request"})
		return
	}

	println("request name: ", request.Name)
	println("request description: ", request.Description)

	file, err := c.FormFile("videoFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	if !filetype.IsVideo(buffer) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is not a valid video"})
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

	go func() {
		if err := utils.HandleVideoProcessing(filename, request); err != nil {
			fmt.Println("Video processing failed for", filename, ":", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully"})
}

func LikeVideoHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - User ID not found"})
		return
	}

	videoID := c.Param("videoId")

	var video models.Video
	if err := config.DB.First(&video, "video_id = ?", videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	action := "liked"

	err := utils.CheckIfVideoIsLikedByUser(userID, videoID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video already liked"})
		return
	}

	video.Likes += 1

	if err := config.DB.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like video"})
		return
	}

	var user models.User
	if err := config.DB.Preload("LikedVideos").First(&user, "user_id = ?", userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	user.LikedVideos = append(user.LikedVideos, &video)
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Video %s successfully", action), "likes": video.Likes})
}

/*

func DislikeVideoHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - User ID not found"})
		return
	}

	videoID := c.Param("videoId")

	var video models.Video
	if err := config.DB.First(&video, "video_id = ?", videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	if video.Likes > 0 {
		video.Likes -= 1
		if err := config.DB.Save(&video).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike video"})
			return
		}
}
} */
