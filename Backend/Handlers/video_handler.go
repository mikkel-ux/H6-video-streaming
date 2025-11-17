/* https://superuser.com/questions/1682333/ffmpeg-how-can-i-get-the-first-frame-of-an-mp4-and-maintain-its-aspect-ratio img from video
// using ffmpeg */

package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	"fmt"
	"net/http"
	"os"
	"strings"

	DTO "VideoStreamingBackend/Models/DTO"
	utils "VideoStreamingBackend/Utils"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	"gorm.io/gorm"
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
	filename = strings.ReplaceAll(filename, " ", "_")
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
	action := "liked"
	var video models.Video
	if err := config.DB.First(&video, "video_id = ?", videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	err := utils.CheckIfVideoIsLikedByUser(userID, videoID)
	if err == nil {
		err := config.DB.Table("user_liked_videos").
			Where("user_user_id = ? AND video_video_id = ?", userID, videoID).
			Delete(nil).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike video"})
			return
		}

		if video.Likes > 0 {
			if err := config.DB.Model(&models.Video{}).
				Where("video_id = ?", videoID).
				UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike video"})
				return
			}
		}
		action = "unliked"
	} else {
		if err := config.DB.Model(&models.Video{}).
			Where("video_id = ?", videoID).
			UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
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
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Video %s successfully", action)})
}

func DislikedVideosHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - User ID not found"})
		return
	}

	videoID := c.Param("videoId")
	action := "disliked"
	var video models.Video
	if err := config.DB.First(&video, "video_id = ?", videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	err := utils.CheckIfVideoIsDislikedByUser(userID, videoID)
	println("error from checkifdisliked:", err)
	if err == nil {
		err := config.DB.Table("user_disliked_videos").
			Where("user_user_id = ? AND video_video_id = ?", userID, videoID).
			Delete(nil).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike video"})
			return
		}

		if video.Dislikes > 0 {
			if err := config.DB.Model(&models.Video{}).
				Where("video_id = ?", videoID).
				UpdateColumn("dislikes", gorm.Expr("dislikes - ?", 1)).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to undislike video"})
				return
			}
		}
		action = "undisliked"
	} else {
		if err := config.DB.Model(&models.Video{}).
			Where("video_id = ?", videoID).
			UpdateColumn("dislikes", gorm.Expr("dislikes + ?", 1)).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike video"})
			return
		}
		var user models.User
		if err := config.DB.Preload("DislikedVideos").First(&user, "user_id = ?", userID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}

		user.DislikedVideos = append(user.DislikedVideos, &video)
		if err := config.DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like video"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Video %s successfully", action)})

}

func GetVideoHandler(c *gin.Context) {
	videoID := c.Param("videoId")
	var video models.Video
	if err := config.DB.First(&video, "video_id = ?", videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	videoResponse := DTO.GetVideoResponse{
		VideoID:     video.VideoID,
		Title:       video.Title,
		Description: video.Description,
		Uploaded:    video.Uploaded,
		URL:         video.URL,
		Thumbnail:   video.Thumbnail,
		Likes:       video.Likes,
		Dislikes:    video.Dislikes,
		Channel:     video.Channel,
		ChannelID:   video.ChannelID,
	}

	c.JSON(http.StatusOK, videoResponse)
}

func VideoStreamHandler(c *gin.Context) {
	videoName := c.Param("videoName")
	fmt.Println("Streaming video from path: ", videoName)
	filePath := "./Uploads/Videos/" + videoName
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read file stats"})
		return
	}

	http.ServeContent(c.Writer, c.Request, videoName, stat.ModTime(), file)
}

func Get30RandomVideosHandler(c *gin.Context) {
	var videos []models.Video
	if err := config.DB.Order("NEWID()").Limit(30).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	var videoPreviews []DTO.VideoPreview
	for _, video := range videos {
		videoPreviews = append(videoPreviews, DTO.VideoPreview{
			VideoID:   video.VideoID,
			Title:     video.Title,
			Thumbnail: video.Thumbnail,
		})
	}

	c.JSON(http.StatusOK, videoPreviews)
}

func GetThumbnailHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetThumbnailHandler reached"})
}
