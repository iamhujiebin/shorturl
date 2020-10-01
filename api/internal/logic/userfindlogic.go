package logic

import (
	"context"
	"shorturl/rpc/userservice/userservice"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserFindLogic {
	return UserFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFindLogic) UserFind(req types.UserFindReq) (*types.UserFindRes, error) {
	resp, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserFindRes{
		UserInfo: types.UserInfo{
			UserId: resp.UserInfo.UserId,
			Name:   resp.UserInfo.Name,
			Age:    resp.UserInfo.Age,
		},
	}, nil
}
