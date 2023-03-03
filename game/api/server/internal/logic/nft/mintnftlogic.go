package nft

import (
	"context"
	"errors"
	"fmt"
	zecreyface "github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	sdk "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/cronjob/zecreyface"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"strings"
)

type MintNftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMintNftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MintNftLogic {
	return &MintNftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MintNftLogic) MintNft(req *types.ReqMintNft) (*types.RespMintNft, error) {
	resp := &types.RespMintNft{}
	player, err := l.getFromMarket(req.AccountName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getFromMarket err=%v", err))
	}
	ok, err := sdk.VerifyMessage(player.Account.AccountPk, req.SignedMessage, req.RawMessage)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("VerifyMessage err=%v", err))
	}
	if !ok {
		return nil, errors.New(fmt.Sprintf("VerifyMessage fail err=%v", err))
	}
	nftInfo, err := l.svcCtx.SdkClient.MintNft(l.svcCtx.Config.CollectionId, req.MediaId, player.Account.AccountName,
		fmt.Sprintf("%s-%s", l.svcCtx.Config.NftPrefix, req.BoxName), //nftName
		req.BoxName, req.BoxId,
		fmt.Sprintf("%s-%s%d", l.svcCtx.Config.NftPrefix, req.BoxName, req.BoxId)) //nftDescription
	if err != nil {
		logx.Error("MintNft", zap.Error(err))
		return nil, errors.New(fmt.Sprintf("MintNft err=%v", err))
	}
	resp.Id = nftInfo.Asset.Id
	return resp, nil
}

func (l *MintNftLogic) getFromMarket(playerName string) (*zecreyface.RespGetAccountByAccountName, error) {
	//delete .zec suffix use zecrey nft sdk not need add .zec suffix will add it automatic
	name := strings.TrimSuffix(playerName, ".zec")
	//set playerPk
	playerInfo, err := sdk.GetAccountInfo(name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAccountInfo err=%v", err))
	}
	return playerInfo, nil
}
