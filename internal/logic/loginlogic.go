package logic

import (
	"context"
	"douyin/internal/models"
	"fmt"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	user := models.User{
		UserName: req.Username,
		Password: req.Password,
	}
	result := l.svcCtx.DbEngin.First(&user)
	fmt.Println(user.UserId)
	return &types.LoginResponse{
		Status_code: 0,
		Status_msg:  "登录成功",
		User_id:     user.UserId,
		Token:       user.UserName + user.Password,
	}, result.Error
}
