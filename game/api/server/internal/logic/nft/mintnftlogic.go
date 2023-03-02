package nft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	sdk "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/cronjob/zecreyface"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/types"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/model"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type MintNftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type gameData struct {
	NftId   int64  `json:"nft_id"`
	BoxId   int64  `json:"box_id"`
	BoxName string `json:"box_name"`
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
	player, err := l.svcCtx.Db.Player.Get(req.AccountName)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New(fmt.Sprintf("1 err=%v", err))
	}
	if err == gorm.ErrRecordNotFound {
		player, err = l.getFromMarket(req.AccountName)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("2 err=%v", err))
		}
		err = l.svcCtx.Db.Player.Update(player)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("3 err=%v", err))
		}
	}
	ok, err := sdk.VerifyMessage(player.L2publicKey, req.SignedMessage, req.RawMessage)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("4 err=%v", err))
	}
	if !ok {
		return nil, errors.New(fmt.Sprintf("5 err=%v", err))
	}
	nftInfo, err := l.svcCtx.SdkClient.MintNft(l.svcCtx.Config.CollectionId, req.MediaId, player.AccountName,
		fmt.Sprintf("%s-%s%d", l.svcCtx.Config.NftPrefix, req.BoxName, req.BoxId),
		fmt.Sprintf("%s%d", req.BoxName, req.BoxId),
		fmt.Sprintf("%s-%s%d", l.svcCtx.Config.NftPrefix, req.BoxName, req.BoxId))
	if err != nil {
		logx.Error("MintNft", zap.Error(err))
		return nil, errors.New(fmt.Sprintf("6 err=%v", err))
	}
	resp.Id = nftInfo.Asset.Id
	var gameDatas []*gameData
	if player.Data != "" {
		err := json.Unmarshal([]byte(player.Data), &gameDatas)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("19 err=%v", err))
		}
	}
	gameDatas = append(gameDatas, &gameData{
		NftId:   nftInfo.Asset.Id,
		BoxId:   req.BoxId,
		BoxName: req.BoxName,
	})
	data, err := json.Marshal(gameDatas)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("19 err=%v", err))
	}
	player.Data = string(data)
	err = l.svcCtx.Db.Player.Update(player)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("3 err=%v", err))
	}
	return resp, nil
}

func (l *MintNftLogic) getFromMarket(playerName string) (*model.Player, error) {
	//delete .zec suffix use zecrey nft sdk not need add .zec suffix will add it automatic
	name := strings.TrimSuffix(playerName, ".zec")
	//set playerPk
	playerInfo, err := sdk.GetAccountInfo(name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("7 err=%v", err))
	}
	player := &model.Player{
		AccountIndex: playerInfo.Account.AccountIndex,
		AccountName:  playerInfo.Account.AccountName,
		L2publicKey:  playerInfo.Account.AccountPk,
	}

	if err := l.svcCtx.Db.Player.Create(player); err != nil {
		return nil, errors.New(fmt.Sprintf(" 8 create player failed err=%v", err))
	}
	return player, nil
}
