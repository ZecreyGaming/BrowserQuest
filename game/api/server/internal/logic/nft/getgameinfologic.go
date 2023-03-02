package nft

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGameInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGameInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGameInfoLogic {
	return &GetGameInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGameInfoLogic) GetGameInfo(req *types.ReqGetGameInfo) (*types.RespGetGameInfo, error) {
	resp := &types.RespGetGameInfo{}
	player, err := l.svcCtx.Db.Player.Get(req.AccountName)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, err
	}
	resp.GameData = player.Data
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.GameData = "[]"
	}
	return resp, nil
}
