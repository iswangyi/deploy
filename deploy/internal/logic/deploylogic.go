package logic

import (
	"context"

	"ccdd/deploy/internal/svc"
	"ccdd/deploy/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeployLogic {
	return DeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeployLogic) Deploy(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
