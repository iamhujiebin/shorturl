package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandHandler2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandHandler2Logic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandHandler2Logic {
	return ExpandHandler2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandHandler2Logic) ExpandHandler2(req types.ExpandReq) (*types.ExpandResp, error) {
	// todo: add your logic here and delete this line

	return &types.ExpandResp{}, nil
}
