package models

import "time"

type Video struct {
	VideoID     int64      `gorm:"primaryKey;type:bigint;autoIncrement" json:"videoId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Uploaded    time.Time  `json:"uploaded"`
	URL         string     `json:"url"`
	Thumbnail   string     `json:"thumbnail"`
	Likes       int64      `json:"likes"`
	Dislikes    int64      `json:"dislikes"`
	ChannelID   int64      `json:"channelId"`
	Channel     *Channel   `json:"channel"`
	Comments    []*Comment `json:"comments"`
}
