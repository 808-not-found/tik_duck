package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
)

type User struct {
	ID            int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	Name          string    `gorm:"column:name;unique;NOT NULL"`
	Password      string    `gorm:"column:password;NOT NULL"`
	Salt          string    `gorm:"column:salt;NOT NULL"`
	FollowCount   int       `gorm:"column:follow_count;default:0;NOT NULL"`
	FollowerCount int       `gorm:"column:follower_count;default:0;NOT NULL"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// 传入批量用户id 返回用户信息
func MGetUsers(ctx context.Context, userIds []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIds) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info.
// 传入用户信息内容 在数据库创建用户
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info.
// 传入用户名称 查找用户信息
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
