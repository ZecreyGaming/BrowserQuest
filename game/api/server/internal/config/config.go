package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Postgres struct {
		DataSource string
	}
	LogConf      logx.LogConf
	AccountName  string
	Seed         string
	NftPrefix    string
	CollectionId int64
}
