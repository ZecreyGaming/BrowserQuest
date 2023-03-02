package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	LogConf      logx.LogConf
	Seed         string
	NftPrefix    string
	CollectionId int64
}
