package models

import "time"

const (
	TableNameVote = "votes"
)

type Vote struct {
	InheritPostId int64     `json:"inheritPostId" gorm:"not null;primaryKey"`
	ExpiredAt     time.Time `json:"expiredAt" gorm:"not null;default:'now()'"`
	IsExpired     bool      `json:"isExpired" gorm:"not null;default:false"`
	Register      bool      `json:"register" gorm:"not null;default:false"`
}
