package handler

import (
	"net/http"

	"ccdd/internal/logic"
	"ccdd/internal/svc"
	"ccdd/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func HelmHadnlerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelmHadnlerLogic(r.Context(), ctx)
		resp, err := l.HelmHadnler(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
