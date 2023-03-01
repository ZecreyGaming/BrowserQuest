package svc

import (
	"github.com/ZecreyGaming/zecrey_treasure_hunt/db"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/config"
	sdk "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/cronjob/zecreyface"
)

type ServiceContext struct {
	Config    config.Config
	Db        *db.Db
	SdkClient *sdk.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := db.NewDb(c.Postgres.DataSource)
	sdkClient, err := sdk.GetClient(c.AccountName, c.Seed, c.NftPrefix, c.CollectionId)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		Db:        db,
		SdkClient: sdkClient,
	}
}
