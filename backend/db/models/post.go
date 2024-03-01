package models

import (
	"database/sql"
	"time"
)

const (
	TableNamePost = "posts"
)

type Post struct {
	Id          int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt   time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt   time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt   sql.NullTime `json:"deletedAt"`
	FromRoleId  int64        `json:"fromRoleId" gorm:"not null"`
	ToSectionId int64        `json:"toSectionId" gorm:"not null"`
	LikesCount  int64        `json:"likesCount" gorm:"not null;default: 0"`
	IsTop       bool         `json:"isTop" gorm:"not null;default: false"`
	IsHighlight bool         `json:"isHighlight" gorm:"not null;default:false"`
	Title       string       `json:"title" gorm:"not null;size:127"`
	Body        string       `json:"body" gorm:"not null;size:127"`
}
