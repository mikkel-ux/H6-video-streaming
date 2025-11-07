package models

type Channel struct {
	ChannelID   int64    `gorm:"primaryKey;autoIncrement" json:"channelId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserID      int64    `gorm:"uniqueIndex" json:"userId"`
	User        *User    `json:"user"`
	Videos      []*Video `json:"videos"`
}
