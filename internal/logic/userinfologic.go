package logic

import (
	"context"
	"douyin/internal/models"
	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	user := models.User{
		UserId: req.User_id,
	}
	result := l.svcCtx.DbEngin.First(&user)
	return &types.UserInfoResponse{
		Status_code: 0,
		Status_msg:  "用户信息返回成功",
		User: types.User(User{
			user.UserId,
			user.UserName,
			user.FollowCount,
			user.FollowerCount,
		}),
	}, result.Error
}

type User struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Follow_count   int    `json:"follow_count"`
	Follower_count int    `json:"follower_count"`
}
