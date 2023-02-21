package pack

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
)

func User(m *db.User) *user.User {
	return &user.User{
		Id:            m.ID,
		Name:          m.Name,
		FollowCount:   &m.FollowCount,
		FollowerCount: &m.FollowerCount,
		IsFollow:      false,
	}
}

// 传入的是 数据库的原始值 这里返回的应该是封装好了 用于 rpc 的值.
func Video(ctx context.Context, m *db.Video, myID int64) (*user.Video, error) {
	var res *user.Video
	if m == nil {
		return res, nil
	}

	// 查询数据库中该视频的作者 db User
	dbAuthor, err := db.GetUser(ctx, m.AuthorID)
	if err != nil {
		return nil, err
	}

	// 封装转换 rpc User
	var rpcAuthor *user.User
	if myID == 0 { // 未登录
		rpcAuthor = User(&dbAuthor)
	} else { // 登录
		rpcAuthor, err = DBUserToRPCUser(&dbAuthor, myID)
		if err != nil {
			return nil, err
		}
	}

	var like bool
	err = db.DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", myID, m.ID).Error
	// err = db.IsFavorite(ctx, myID, m.ID)
	if err != nil {
		like = false
	} else {
		like = true
	}
	// TODO: 完成以下
	return &user.Video{
		Id:            m.ID,
		Author:        rpcAuthor,
		PlayPath:      "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + m.FilePath,
		CoverPath:     "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + m.CoverPath,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    like,
		Title:         m.Title,
	}, nil
}

func Videos(ctx context.Context, ms []*db.Video, myID int64) ([]*user.Video, error) {
	videos := make([]*user.Video, 0)
	for _, m := range ms {
		n, err := Video(ctx, m, myID)
		if err != nil {
			return nil, err
		}
		videos = append(videos, n)
	}
	return videos, nil
}
