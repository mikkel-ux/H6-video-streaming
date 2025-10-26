package models

type User struct {
	UserID        int64          `gorm:"primaryKey;type:bigint;autoIncrement" json:"userId"`
	Name          string         `json:"name"`
	Age           int            `json:"age"`
	Email         string         `gorm:"unique" json:"email"`
	Password      string         `json:"password"`
	Channels      *Channel       `json:"channels"`
	RefreshTokens []RefreshToken `json:"refreshTokens"`
}
