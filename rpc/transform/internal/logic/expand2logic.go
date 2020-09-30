package logic

import (
	"context"

	"shorturl/rpc/transform/internal/svc"
	transform "shorturl/rpc/transform/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type Expand2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpand2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Expand2Logic {
	return &Expand2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Expand2Logic) Expand2(in *transform.Expand2Req) (*transform.Expand2Resp, error) {
	res, err := l.svcCtx.MongoModel.FindOne(in.Shorten)
	if err != nil {
		return nil, err
	}
	return &transform.Expand2Resp{
		Url: res.Url,
	}, nil
}
