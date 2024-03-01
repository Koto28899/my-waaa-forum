package models

import (
	"database/sql"
	"time"
)

const (
	TableNameReply = "replies"
)

type Reply struct {
	Id          int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt   time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt   time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt   sql.NullTime `json:"deletedAt"`
	FromRoleId  int64        `json:"fromRoleId" gorm:"not null"`
	ToCommentId int64        `json:"toCommentId" gorm:"not null"`
	LikesCount  bool         `json:"likesCount" gorm:"not null;default: 0"`
	Body        string       `json:"body" gorm:"not null;size: 1023"`
}
