// Code generated by goctl. DO NOT EDIT.
package types

type ReqMintNft struct {
	RawMessage    string `form:"raw_message"`
	SignedMessage string `form:"signed_message"`
	AccountName   string `form:"account_name"`
	MediaId       string `form:"media_id"`
	BoxId         int64  `form:"box_id"`
	BoxName       string `form:"box_name"` //Sword of Valour-Defender 1 2 3 4
}

type RespMintNft struct {
	Id int64 `json:"id"`
}
