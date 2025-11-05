package models

import "time"

type Video struct {
	VideoID     int64      `gorm:"primaryKey;autoIncrement" json:"videoId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Uploaded    time.Time  `json:"uploaded"`
	URL         string     `json:"url"`
	Thumbnail   string     `json:"thumbnail"`
	Likes       int64      `json:"likes" gorm:"default:0"`
	Dislikes    int64      `json:"dislikes" gorm:"default:0"`
	ChannelID   *int64     `json:"channelId"` /* TODO change this back to can't be null when we can make a channle */
	Channel     *Channel   `json:"channel"`
	Comments    []*Comment `json:"comments"`
}
