package logic

import (
	"context"

	"shorturl/rpc/userservice/internal/svc"
	user_service "shorturl/rpc/userservice/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user_service.GetUserInfoReq) (*user_service.GetUserInfoRes, error) {
	res, err := l.svcCtx.UserModel.FindOne(int(in.UserId))
	if err != nil {
		return nil, err
	}
	return &user_service.GetUserInfoRes{
		UserInfo: &user_service.UserInfo{
			UserId: int64(res.UserId),
			Name:   res.Nickname,
			Age:    int32(res.Age),
		},
	}, nil
}
