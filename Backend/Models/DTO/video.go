package dto

import (
	"time"
)

type UploadVideoRequest struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	ChannelID   *int64 `form:"channelId" binding:"required"`
}

type GetVideoResponse struct {
	VideoID     int64           `json:"videoId"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Uploaded    time.Time       `json:"uploaded"`
	URL         string          `json:"url"`
	Thumbnail   string          `json:"thumbnail"`
	Likes       int64           `json:"likes"`
	Dislikes    int64           `json:"dislikes"`
	Channel     *ChannelSummary `json:"channel"`
	ChannelID   *int64          `json:"channelId"`
	IsLiked     bool            `json:"isLiked"`
	IsDisliked  bool            `json:"isDisliked"`
}

type VideoPreview struct {
	VideoID   int64  `json:"videoId"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}
