package models

import (
	"database/sql"
	"time"
)

const (
	TableNameNotification = "notifications"

	NotificationTypeReply     = "reply"
	NotificationTypeLike      = "like"
	NotificationTypeAt        = "at"
	NotificationTypeFromAdmin = "from_admin"

	SecneTypePost    = "post"
	SecneTypeComment = "comment"
	SecneTypeReply   = "reply"
)

type Notification struct {
	Id               int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt        time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt        time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt        sql.NullTime `json:"deletedAt"`
	NotificationType string       `json:"notificationType" gorm:"not null"`
	FromRoleId       int64        `json:"fromRoleId" gorm:"not null"`
	SceneType        string       `json:"sceneType" gorm:"not null"`
	SceneId          int64        `json:"sceneId" gorm:"not null"`
	ToRoleId         int64        `json:"toRoleId" gorm:"not null"`
}
