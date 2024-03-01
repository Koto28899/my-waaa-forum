package models

import (
	"database/sql"
	"time"
)

const (
	TableNameVoteOption = "vote_options"
)

type VoteOption struct {
	Id        int64        `json:"id" gorm:"<-false;default:(-);primaryKey;not null"`
	CreatedAt time.Time    `json:"createdAt" gorm:"<-false;default:'now()';not null"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"default:'now()'; not null"`
	DeletedAt sql.NullTime `json:"deletedAt"`
	VoteId    int64        `json:"voteId" gorm:"not null"`
	Info      string       `json:"info" gorm:"not null;size:127"`
	Number    int64        `json:"number" gorm:"not null;default:0"`
}
