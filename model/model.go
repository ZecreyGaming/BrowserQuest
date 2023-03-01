package model

import (
	"gorm.io/gorm"
	"time"
)

type Player struct {
	gorm.Model
	AccountIndex int64  `gorm:"index" json:"account_index"`
	AccountName  string `json:"account_name" gorm:"index;not null;size:128"`
	L2publicKey  string `json:"l2public_key"`
	NftId        int64  `json:"nft_id"`
}

type Game struct {
	gorm.Model
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
