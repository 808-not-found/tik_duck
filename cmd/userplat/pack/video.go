package pack

import (
	"context"
	"errors"
	"time"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	allerrors "github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type Like struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LikeTime time.Time `gorm:"column:like_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;NOT NULL"`
	VideoID  int64     `gorm:"column:video_id;NOT NULL"`
}

func (m *Like) TableName() string {
	return "like"
}

type Follow struct {
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
}

func (m *Follow) TableName() string {
	return "follow"
}

func DBUserToRPCUser(m *db.User, fromID int64) (*userplat.User, error) {
	if m == nil {
		return nil, allerrors.ErrDBUserToRPCUserVali()
	}
	var IsFollowShip bool
	var reserr error
	f := Follow{}

	err := db.DB.Where("from_user_id = ? AND to_user_id = ?", fromID, m.ID).Find(&f).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		IsFollowShip = false
	case err != nil:
		return nil, allerrors.ErrDBUserToRPCUserRun()
	default:
		IsFollowShip = true
	}
	return &userplat.User{
		Id:            m.ID,
		Name:          m.Name,
		FollowCount:   &m.FollowCount,
		FollowerCount: &m.FollowerCount,
		IsFollow:      IsFollowShip,
	}, reserr
}

func User(m *db.User) *userplat.User {
	return &userplat.User{
		Id:            m.ID,
		Name:          m.Name,
		FollowCount:   &m.FollowCount,
		FollowerCount: &m.FollowerCount,
		IsFollow:      false,
	}
}

// 传入的是 数据库的原始值 这里返回的应该是封装好了 用于 rpc 的值.
func Video(ctx context.Context, m *db.Video, myID int64) (*userplat.Video, error) {
	var res *userplat.Video
	if m == nil {
		return res, nil
	}

	// 查询数据库中该视频的作者 db User
	dbAuthor, err := db.GetUser(ctx, m.AuthorID)
	if err != nil {
		return nil, err
	}

	// 封装转换 rpc User
	var rpcAuthor *userplat.User
	if myID == 0 { // 未登录
		rpcAuthor = User(&dbAuthor)
	} else { // 登录
		rpcAuthor, err = DBUserToRPCUser(&dbAuthor, myID)
		if err != nil {
			return nil, err
		}
	}
	// TODO: 完成以下
	return &userplat.Video{
		Id:            m.ID,
		Author:        rpcAuthor,
		PlayUrl:       "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + m.FilePath,
		CoverUrl:      m.CoverPath,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    true, // 残缺
		Title:         m.Title,
	}, nil
}

func Videos(ctx context.Context, ms []*db.Video, myID int64) ([]*userplat.Video, error) {
	videos := make([]*userplat.Video, 0)
	for _, m := range ms {
		n, err := Video(ctx, m, myID)
		if err != nil {
			return nil, err
		}
		videos = append(videos, n)
	}
	return videos, nil
}
