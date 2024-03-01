package models

import (
	"database/sql"
	"time"
)

const (
	TableNameVoteEvent = "vote_events"
)

type VoteEvent struct {
	Id           int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt    time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt    time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt    sql.NullTime `json:"deletedAt"`
	FromRoleId   int64        `json:"fromRoleId" gorm:"not null"`
	ToVoteId     int64        `json:"toVoteId" gorm:"not null"`
	ToVoteOption int64        `json:"toVoteOption" gorm:"not null"`
}
