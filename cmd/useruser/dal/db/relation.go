package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	Name          string    `gorm:"column:name;unique;NOT NULL"`
	Password      string    `gorm:"column:password;NOT NULL"`
	Salt          string    `gorm:"column:salt;NOT NULL"`
	FollowCount   int64     `gorm:"column:follow_count;default:0;NOT NULL"`
	FollowerCount int64     `gorm:"column:follower_count;default:0;NOT NULL"`
}

type Follow struct {
	gorm.Model
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCERMENT"`
	FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

func (u *Follow) TableName() string {
	return consts.FollowTableName
}

// myID 关注 toID.
func FollowAction(ctx context.Context, myID int64, toID int64) error {
	// 增加follow count
	var myUser *User
	conn := DB.WithContext(ctx).Where("id = ?", myID).First(&myUser).Update("id", myUser.FollowCount+1)
	if err := conn.Error; err != nil {
		return err
	}
	// 增加follower count
	var toUser *User
	conn = DB.WithContext(ctx).Where("id = ?", toID).First(&toUser).Update("id", toUser.FollowerCount+1)
	if err := conn.Error; err != nil {
		return err
	}
	// 增加一条记录到follow表
	follow := Follow{
		FromUserID: myID,
		ToUserID:   toID,
	}
	conn = DB.WithContext(ctx).Create(follow)
	if err := conn.Error; err != nil {
		return err
	}
	return nil
}

func UnFollowAction(ctx context.Context, myID int64, toID int64) error {
	// 减少follow count
	var myUser *User
	conn := DB.WithContext(ctx).Where("id = ?", myID).Find(&myUser).Update("follow_count", myUser.FollowCount-1)
	if err := conn.Error; err != nil {
		return err
	}
	// 减少follower count
	var toUser *User
	conn = DB.WithContext(ctx).Where("id = ?", toID).First(&toUser).Update("id", toUser.FollowerCount-1)
	if err := conn.Error; err != nil {
		return err
	}
	// 删除follow表中的一条记录
	follow := Follow{
		FromUserID: myID,
		ToUserID:   toID,
	}
	conn = DB.WithContext(ctx).Delete(follow)
	if err := conn.Error; err != nil {
		return err
	}

	return nil
}

func GetFollowList(ctx context.Context, myID int64) ([]*User, error) {
	var res []*User
	return res, nil
}

func GetFollowerList(ctx context.Context, myID int64) ([]*User, error) {
	var res []*User
	return res, nil
}

func GetFriendList(ctx context.Context, myID int64) ([]*User, error) {
	var res []*User
	return res, nil
}