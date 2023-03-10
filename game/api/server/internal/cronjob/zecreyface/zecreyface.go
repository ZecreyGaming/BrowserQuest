package zecreyface

import (
	"encoding/json"
	"fmt"
	zecreyface "github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	legendSdk "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"
	"time"
)

type Client struct {
	z            *zecreyface.Client
	nftPrefix    string
	acccountName string
	seed         string
	l2pk         string
	collectionId int64
}

func GetClient(accountName, seed, nftPrefix string, collectionId int64) (*Client, error) {
	z, err := zecreyface.NewClient(accountName, seed)
	if err != nil {
		return nil, err
	}
	_, l2pk, seed := z.GetMyInfo()
	//Id, err := zecreyface.GetDefaultCollectionId(accountName)
	//if err != nil {
	//	return nil, err
	//}
	return &Client{
		z:            z,
		acccountName: accountName,
		seed:         seed,
		l2pk:         l2pk,
		collectionId: collectionId,
		nftPrefix:    nftPrefix}, nil
}

func (c *Client) MintNft(collectionId int64, MediaId string, toAccountName string, nftName string, boxName string, boxId int64, nftDescription string) (*zecreyface.RespCreateAsset, error) {
	assetStats := zecreyface.Stat{
		Name:     "time",
		Value:    time.Now().UnixMilli(),
		MaxValue: 0,
	}
	assetLevel := zecreyface.Level{
		Name:     "boxId",
		Value:    boxId,
		MaxValue: 0,
	}
	assetProperty := zecreyface.Propertie{
		Name:  "boxName",
		Value: boxName,
	}
	// get content hash

	_Stats := []zecreyface.Stat{assetStats}
	_Property := []zecreyface.Propertie{assetProperty}
	_Level := []zecreyface.Level{assetLevel}
	_LevelByte, err := json.Marshal(_Level)
	if err != nil {
		return nil, err
	}
	_StatsByte, err := json.Marshal(_Stats)
	if err != nil {
		return nil, err
	}
	_PropertyByte, err := json.Marshal(_Property)
	if err != nil {
		return nil, err
	}
	nftInfo, err := c.z.MintNft(collectionId, toAccountName,
		fmt.Sprintf("https://res.cloudinary.com/zecrey/image/upload/%s", MediaId), nftName,
		nftDescription, MediaId,
		string(_PropertyByte), string(_LevelByte), string(_StatsByte))
	if err != nil {
		return nil, err
	}
	return nftInfo, nil
}

func (c *Client) SignMessage(message string) (string, error) {
	_, _, seed := c.z.GetMyInfo()
	return zecreyface.SignMessage(seed, message)
}

func VerifyMessage(l2publicKey, eddsaSig, rawMessage string) (bool, error) {
	return zecreyface.VerifyMessage(l2publicKey, eddsaSig, rawMessage)
}

func (c *Client) GetAccountWinNfts(collectionId int64, accountName string) ([]*zecreyface.HauaraNftInfo, error) {
	result, err := zecreyface.GetCollectionAccountNftsByIregex(collectionId, accountName, c.nftPrefix)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetCollectionWinNfts(collectionId int64) ([]*zecreyface.HauaraNftInfo, error) {
	result, err := zecreyface.GetCollectionNftsByIregex(collectionId, c.nftPrefix)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAccountInfo(accountName string) (*zecreyface.RespGetAccountByAccountName, error) {
	return zecreyface.GetAccountByAccountName(accountName)
}

func GetAccountInfoBySeed(privateKey string) (*legendSdk.RespGetAccountInfoByPubKey, string, error) {
	_, seed, err := zecreyface.GetSeedAndL2Pk(privateKey)
	if err != nil {
		return nil, "", fmt.Errorf("GetSeedAndL2Pk err=%s", err.Error())
	}
	accountInfo, err := zecreyface.GetAccountInfoBySeed(seed[2:])

	return accountInfo, seed[2:], err
}
