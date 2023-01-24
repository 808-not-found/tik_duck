package pack

import (
	"errors"
	"fmt"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"gorm.io/gorm"
)

var ErrSQL = errors.New("查询失败")
var ErrVali = errors.New("传入查询对象为空")

func ErrDBUserToRPCUserSQL() error {
	return fmt.Errorf("ErrDBUserToRpcUser %w", ErrSQL)
}
func ErrDBUserToRPCUserVali() error {
	return fmt.Errorf("ErrDBUserToRpcUser %w", ErrVali)
}

type Follow struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int       `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int       `gorm:"column:to_user_id;NOT NULL"`
}

func (m *Follow) TableName() string {
	return "follow"
}

func DBUserToRPCUser(m *db.User, fromID int64, toid int64) (*user.User, error) {
	if m == nil {
		return nil, ErrDBUserToRPCUserVali()
	}
	var IsFollowShip bool
	var reserr error
	f := Follow{}

	err := db.DB.Where("FromUserID = ? AND ToUserID = ?", fromID, toid).Find(&f).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		IsFollowShip = false
	case err != nil:
		return nil, ErrDBUserToRPCUserSQL()
	default:
		IsFollowShip = true
	}
	return &user.User{
		//Id:   int64(m.ID),//ToDo:等待类型更新
		Name: m.Name,
		// FollowCount:   int64(m.FollowCount),//ToDo:等待类型更新
		// FollowerCount: int64(m.FollowerCount),//ToDo:等待类型更新
		IsFollow: IsFollowShip,
	}, reserr
}

func DBUsersToRPCUsers(ms []*db.User, fromID int64, toids []int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	var reserr error
	for i, m := range ms {
		userinfo, err := DBUserToRPCUser(m, fromID, toids[i])
		if err != nil {
			reserr = err
		}
		users = append(users, userinfo)
	}
	return users, reserr
}
