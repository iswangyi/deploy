package handler

import (
	"net/http"

	"ccdd/deploy/internal/logic"
	"ccdd/deploy/internal/svc"
	"ccdd/deploy/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func DeployHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeployLogic(r.Context(), ctx)
		resp, err := l.Deploy(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
