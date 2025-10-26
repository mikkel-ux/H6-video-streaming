package models

import "time"

type RefreshToken struct {
	Token     string     `json:"token" gorm:"primaryKey;type:varchar(500)"`
	UserID    int64      `json:"userID" gorm:"not null;column:userID;"`
	ExpiresAt time.Time  `json:"expiresAt" gorm:"not null;index:idx_expires_at;column:expiresAt;"`
	RevokedAt *time.Time `json:"revokedAt" gorm:"index:idx_revoked_at;column:revokedAt;"`
	Revoked   bool       `json:"revoked" gorm:"not null;default:false;index:idx_revoked;column:revoked;"`
	User      User       `json:"user"`
}

func (RefreshToken) TableName() string {
	return "refreshTokens"
}
