package models

import (
	"database/sql"
	"time"
)

const (
	TableNameSection = "section"
)

type Section struct {
	Id          int64          `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt   sql.NullTime   `json:"deletedAt"`
	SectionName string         `json:"sectionName" gorm:"not null;size:16"`
	Statement   sql.NullString `json:"statement" gorm:"size:127"`
	ManagerId   int64          `json:"managerId" gorm:"not null"`
}
