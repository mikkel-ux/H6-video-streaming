/* https://superuser.com/questions/1682333/ffmpeg-how-can-i-get-the-first-frame-of-an-mp4-and-maintain-its-aspect-ratio img from video
// using ffmpeg */

package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
)

type UploadVideoRequest struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
}

// TODO implement video processing (e.g., change metadata so it's suitable for streaming, thumbnail generation)
func handleVideoProcessing(tempPath string, videoDetails UploadVideoRequest) error {
	println("Processing video:", tempPath)
	thumbnailPath := strings.Replace(tempPath, "TempVideoPath", "Images", 1)
	uploadPath := strings.Replace(tempPath, "TempVideoPath", "Videos", 1)

	lastDot := strings.LastIndex(thumbnailPath, ".")
	if lastDot != -1 {
		thumbnailPath = thumbnailPath[:lastDot] + "___thumbnail.jpg"
		println(thumbnailPath)
	} else {
		thumbnailPath = thumbnailPath + "___thumbnail.jpg"
	}
	/* command er fra en på stackExchange men jeg har selv sat den op så go kan køre den */
	cmd := exec.Command("ffmpeg", "-i", tempPath, "-vf", "scale=iw*sar:ih,setsar=1", "-vframes", "1", thumbnailPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error processing video:", err)
		return err
	}

	cmd = exec.Command("ffmpeg", "-i", tempPath,
		"-c:v", "copy", "-c:a", "copy", "-movflags", "+faststart",
		uploadPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error moving file to uploads: %v", err)
	}
	os.Remove(tempPath)

	if err := config.DB.Create(&models.Video{
		Title:       videoDetails.Name,
		Description: videoDetails.Description,
		URL:         uploadPath,
		Thumbnail:   thumbnailPath,
		Uploaded:    time.Now(),
	}).Error; err != nil {
		return fmt.Errorf("error saving video to database: %v", err)
	}
	return nil
}

func UploadVideoHandler(c *gin.Context) {
	var request UploadVideoRequest

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
		if err := handleVideoProcessing(filename, request); err != nil {
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
	video.Likes += 1
	if err := config.DB.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like video"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	/* user.LikedVideos = append(user.LikedVideos, video)
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user liked videos"})
		return
	} */

	c.JSON(http.StatusOK, gin.H{"message": "Video liked successfully", "likes": video.Likes})
}
