package models

type Comment struct {
	CommentID int64      `gorm:"primaryKey;type:bigint;autoIncrement" json:"commentId"`
	Content   string     `json:"content"`
	UserID    int64      `json:"userId"`
	VideoID   int64      `json:"videoId"`
	Video     *Video     `json:"video"`
	User      *User      `json:"user"`
	CreatedAt int64      `json:"createdAt"`
	ParentID  *int64     `json:"parentId,omitempty"`
	Parent    *Comment   `json:"parent,omitempty"`
	Replies   []*Comment `gorm:"foreignKey:ParentID" json:"replies"`
}
