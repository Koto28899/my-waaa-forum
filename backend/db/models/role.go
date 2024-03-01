package models

import (
	"database/sql"
	"time"
)

const (
	TableNameRole = "roles"

	RoleTypeUser  = "user"
	RoleTypeAdmin = "admin"
)

type Role struct {
	Id                int64          `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt         time.Time      `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt         time.Time      `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt         sql.NullTime   `json:"deletedAt"`
	Email             string         `json:"email" gorm:"size:321;unique;not null"`
	HashedPwd         string         `json:"hashedPwd" gorm:"not null"`
	PasswordChangedAt time.Time      `json:"passwordChangedAt" gorm:"default:'now()';not null"`
	RoleName          string         `json:"roleName" gorm:"size:16; not null"`
	Type              string         `json:"type" gorm:"not null;default:'user'"`
	Statement         sql.NullString `json:"statement" gorm:"size:127"`
	PostsCount        int64          `json:"postsCounts" gorm:"default: 0;not null"`
	CommentsCount     int64          `json:"commentsCounts" gorm:"default: 0;not null"`
	RepliesCount      int64          `json:"repliesCounts" gorm:"default: 0;not null"`
	LikesCount        int64          `json:"likesCounts" gorm:"default: 0;not null"`
}
