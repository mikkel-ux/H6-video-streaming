package models

type User struct {
	UserID        int64           `gorm:"primaryKey;autoIncrement" json:"userId"`
	FirstName     string          `json:"firstName"`
	LastName      string          `json:"lastName"`
	UserName      string          `json:"userName"`
	Age           int             `json:"age"`
	Email         string          `gorm:"unique" json:"email"`
	Password      string          `json:"password"`
	Channels      *Channel        `json:"channels"`
	RefreshTokens []*RefreshToken `json:"refreshTokens"`
	VideoHistory  []*Video        `gorm:"many2many:user_video_history;" json:"videoHistory"`
	WatchLater    []*Video        `gorm:"many2many:user_watch_later;" json:"watchLater"`
}
