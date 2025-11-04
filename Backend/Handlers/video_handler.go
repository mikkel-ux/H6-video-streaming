/* https://superuser.com/questions/1682333/ffmpeg-how-can-i-get-the-first-frame-of-an-mp4-and-maintain-its-aspect-ratio img from video
// using ffmpeg */

package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO implement video processing (e.g., change metadata so it's suitable for streaming, thumbnail generation)
func handleVideoProcessing(tempPath string) error {
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
	return nil
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

	go func() {
		if err := handleVideoProcessing(filename); err != nil {
			fmt.Println("Video processing failed for", filename, ":", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully"})
}
