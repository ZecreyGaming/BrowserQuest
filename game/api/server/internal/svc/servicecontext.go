package svc

import (
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/config"
	sdk "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/cronjob/zecreyface"

	"strings"
)

type ServiceContext struct {
	Config    config.Config
	SdkClient *sdk.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	AccountInfo, err := sdk.GetAccountInfoBySeed(c.Seed)
	if err != nil {
		panic(err)
	}
	sdkClient, err := sdk.GetClient(strings.TrimSuffix(AccountInfo.AccountName, ".zec"), c.Seed, c.NftPrefix, c.CollectionId)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		SdkClient: sdkClient,
	}
}
