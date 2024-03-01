package models

import (
	"database/sql"
	"time"
)

const (
	TableNameComment = "comments"
)

type Comment struct {
	Id         int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt  time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt  time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt  sql.NullTime `json:"deletedAt"`
	FromRoleId int64        `json:"fromRoleId" gorm:"not null"`
	ToPostId   int64        `json:"toPostId" gorm:"not null"`
	LikesCount int64        `json:"likesCount" gorm:"not null;default: 0"`
	IsTop      bool         `json:"isTop" gorm:"not null;default: false"`
	Body       string       `json:"body" gorm:"not null;size: 1023"`
}
