package logic

import (
	"context"
	"douyin/internal/models"
	"fmt"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	user := models.User{
		UserName:      req.Username,
		Password:      req.Password,
		FollowCount:   0,
		FollowerCount: 0,
	}

	result := l.svcCtx.DbEngin.Create(&user)

	//获取插入记录的Id
	var id []int64
	l.svcCtx.DbEngin.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	fmt.Println(id[0])
	fmt.Println(user.UserName)
	fmt.Println(user.Password)
	return &types.RegisterResponse{
		Status_code: 0,
		Status_msg:  "注册成功",
		User_id:     id[0],
		Token:       user.UserName + user.Password,
	}, result.Error
}
