package utils

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	DTO "VideoStreamingBackend/Models/DTO"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO implement video processing (e.g., change metadata so it's suitable for streaming, thumbnail generation)
func HandleVideoProcessing(tempPath string, videoDetails DTO.UploadVideoRequest) error {
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
		ChannelID:   videoDetails.ChannelID,
	}).Error; err != nil {
		return fmt.Errorf("error saving video to database: %v", err)
	}
	return nil
}

func CheckIfVideoIsLikedByUser(userID any, videoID string) error {
	var count int64

	err := config.DB.Table("user_liked_videos").
		Where("user_user_id = ? AND video_video_id = ?", userID, videoID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	} else {
		return fmt.Errorf("like not found")
	}
}

func CheckIfVideoIsDislikedByUser(userID any, videoID string) error {
	var count int64

	err := config.DB.Table("user_disliked_videos").
		Where("user_user_id = ? AND video_video_id = ?", userID, videoID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	} else {
		return fmt.Errorf("dislike not found")
	}
}
