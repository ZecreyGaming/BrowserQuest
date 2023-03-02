package model

type Player struct {
	AccountIndex int64  `gorm:"index" json:"account_index"`
	AccountName  string `json:"account_name" gorm:"index;not null;size:128"`
	L2publicKey  string `json:"l2public_key"`
}
