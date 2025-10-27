package models

import "time"

type RefreshToken struct {
	Token     string     `json:"token" gorm:"primaryKey;type:varchar(500)"`
	UserID    int64      `json:"userID"`
	ExpiresAt time.Time  `json:"expiresAt" gorm:"index:idx_expires_at"`
	RevokedAt *time.Time `json:"revokedAt" gorm:"index:idx_revoked_at"`
	Revoked   bool       `json:"revoked" gorm:"default:false;index:idx_revoked;column:revoked;"`
	User      User       `json:"user"`
}

/* func (RefreshToken) TableName() string {
	return "refreshTokens"
} */
