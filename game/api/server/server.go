package main

import (
	"flag"
	"fmt"

	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/config"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/handler"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f",
	"/etc/config/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	select {}
}
