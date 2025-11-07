package handlers

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	DTO "VideoStreamingBackend/Models/DTO"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChannelHandler(g *gin.Context) {
	channelId := g.Param("channelId")
	var channel models.Channel
	if err := config.DB.Preload("User").Preload("Videos").First(&channel, channelId).Error; err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Channel not found"})
		return
	}

	response := DTO.GetChannelResponse{
		ChannelID:   channel.ChannelID,
		Name:        channel.Name,
		Description: channel.Description,
		User: DTO.UserSummary{
			UserID:   channel.User.UserID,
			UserName: channel.User.UserName,
		},
	}

	for _, video := range channel.Videos {
		response.Videos = append(response.Videos, DTO.VideoPreview{
			VideoID:   video.VideoID,
			Title:     video.Title,
			Thumbnail: video.Thumbnail,
		})
	}

	g.JSON(http.StatusOK, response)
}
