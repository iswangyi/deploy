package logic

import (
	"context"

	"ccdd/internal/svc"
	"ccdd/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelmHadnlerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelmHadnlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelmHadnlerLogic {
	return HelmHadnlerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelmHadnlerLogic) HelmHadnler(req types.HelmReq) (resp *types.HelmResp, err error) {

	return &types.HelmResp{Helm: "12333333"}, nil
}
