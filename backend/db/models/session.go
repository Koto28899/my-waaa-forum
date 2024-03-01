package models

import (
	"time"
)

const (
	TableNameSession = "sessions"
)

type Session struct {
	Id        string    `json:"id" gorm:"<-false;primaryKey;default:(-)"`
	RoleId    int64     `json:"roleId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"<-false;not null;default:'now()'"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null;default:'now()'"`
	UserAgent string    `json:"userAgent" gorm:"not null"`
	ClientIp  string    `json:"clientIp" gorm:"not null"`
	IsBlocked bool      `json:"isBlocked" gorm:"not null;deafult:false"`
}
