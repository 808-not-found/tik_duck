package db

import (
	"context"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// UserId int64  `json:user_id`
	// Token string `json:token`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

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
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info.
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
