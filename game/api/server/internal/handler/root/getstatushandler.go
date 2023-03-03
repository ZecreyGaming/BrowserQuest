package root

import (
	"net/http"

	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/logic/root"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetStatus
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := root.NewGetStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
