package logic

import (
	"GaussConsul/service/user/model"
	"context"
	"errors"
	"google.golang.org/grpc/status"

	"GaussConsul/service/user/rpc/internal/svc"
	"GaussConsul/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name.String,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
