package db

import (
	"context"
	"log"
	"sort"
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
	// conn := DB.Model(&myUser).WithContext(ctx).Where("id = ?", myID)
	conn := DB.WithContext(ctx).Where("id = ?", myID).First(&myUser).Update("follow_count", myUser.FollowCount+1)
	// conn := DB.WithContext(ctx)
	// conn.Where("id = ?", myID).First(&myUser)
	// cnt1 := myUser.FollowCount
	// conn.Model(&myUser).Update("id", cnt1 + 1)
	if err := conn.Error; err != nil {
		return err
	}
	// 增加follower count
	var toUser *User
	conn = DB.WithContext(ctx).Where("id = ?", toID).First(&toUser).Update("follower_count", toUser.FollowerCount+1)
	// conn = DB.WithContext(ctx)
	// conn.Where("id = ?", toID).First(&toUser)
	// cnt2 := myUser.FollowerCount
	// conn.Model(&toUser).Update("id", cnt + 1)
	if err := conn.Error; err != nil {
		return err
	}
	// 增加一条记录到follow表
	follow := Follow{
		FromUserID: myID,
		ToUserID:   toID,
	}
	conn = DB.WithContext(ctx).Create(&follow)
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
	conn = DB.WithContext(ctx).Where("id = ?", toID).First(&toUser).Update("follower_count", toUser.FollowerCount-1)
	if err := conn.Error; err != nil {
		return err
	}
	// 删除follow表中的一条记录
	var follow *Follow
	log.Println("follow结构体", follow)
	conn = DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ?", myID, toID).Delete(&follow)
	if err := conn.Error; err != nil {
		return err
	}

	return nil
}

func GetFollowList(ctx context.Context, myID int64) ([]*User, error) {
	var res []*User
	// 找到所有和 myID 相关的记录
	var followList []*Follow
	conn := DB.WithContext(ctx).Where("from_user_id = ?", myID).Find(&followList)
	if err := conn.Error; err != nil {
		return res, err
	}

	// 获取所有的用户 ID
	setList := make(map[int64]bool)
	for _, value := range followList {
		if value != nil {
			setList[value.ToUserID] = true // nolint:all
		}
	}
	var followIDList []int64
	for k := range setList {
		followIDList = append(followIDList, k)
	}

	// 找到所有对应的用户结构体
	conn = DB.WithContext(ctx).Where("id = ?", followIDList).Find(&res)
	if err := conn.Error; err != nil {
		return res, err
	}

	return res, nil
}

func GetFollowerList(ctx context.Context, userID int64) ([]*User, error) {
	var res []*User
	// 找到所有相关结构体
	var followerList []*Follow
	conn := DB.WithContext(ctx).Where("to_user_id = ?", userID).Find(&followerList)
	if err := conn.Error; err != nil {
		return res, nil
	}
	log.Println("followerList结构体长度：", len(followerList))

	// 获取用户 ID
	setList := make(map[int64]bool)
	for _, value := range followerList {
		if value != nil {
			setList[value.FromUserID] = true
		}
	}
	var followerIDList []int64
	for k := range setList {
		followerIDList = append(followerIDList, k)
	}

	// 找到所有对应的用户结构体
	conn = DB.WithContext(ctx).Where("id = ?", followerIDList).Find(&res)
	if err := conn.Error; err != nil {
		return res, err
	}

	return res, nil
}

func GetFriendList(ctx context.Context, myID int64) ([]*User, error) {
	var res []*User
	myFollow, err := GetFollowList(ctx, myID)
	if err != nil {
		return res, nil
	}
	myFollower, err := GetFollowerList(ctx, myID)
	if err != nil {
		return res, nil
	}
	// 排序两个数组
	sort.Slice(myFollow, func(i, j int) bool {
		return myFollow[i].ID < myFollow[j].ID
	})
	sort.Slice(myFollower, func(i, j int) bool {
		return myFollower[i].ID < myFollower[j].ID
	})

	// 双指针取出重复值
	n, m := len(myFollow), len(myFollower)
	for i, j := 0, 0; i < n && j < m; i++ {
		for myFollower[j].ID < myFollow[i].ID {
			j++
		}
		if myFollow[i].ID == myFollower[j].ID {
			res = append(res, myFollow[i])
			j++
		}
	}

	return res, nil
}
