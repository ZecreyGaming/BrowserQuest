package root

import (
	"context"

	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatusLogic {
	return &GetStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatusLogic) GetStatus(req *types.ReqGetStatus) (resp *types.RespGetStatus, err error) {
	resp = &types.RespGetStatus{ServerVersion: "0.0.1"}
	return resp, nil
}
