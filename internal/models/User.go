package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	UserId        int64  `json:"user_id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

//在创建前检验验证一下密码的有效性
func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Password) < 6 {
		return errors.New("密码太简单了")
	}
	//对密码进行加密存储
	//u.Passwd = utils.Password(u.Passwd)
	return nil
}
