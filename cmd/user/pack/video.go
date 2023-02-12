package pack

import (
	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
)

// 传入的是 数据库的原始值 这里返回的应该是封装好了 用于 rpc 的值.
func Video(m *db.Video) *user.Video {
	if m == nil {
		return nil
	}
	// TODO: 完成以下
	return &user.Video{
		Id:            m.ID,
		Author:        &user.User{Id: 1, Name: "FakeUser", FollowCount: nil, FollowerCount: nil, IsFollow: false}, // 残缺
		PlayPath:      "http://" + consts.WebServerPublicIP + ":" + consts.WebServerPort + "/" + m.FilePath,
		CoverPath:     m.CoverPath,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    false, // 残缺
		Title:         m.Title,
	}
}

func Videos(ms []*db.Video) []*user.Video {
	videos := make([]*user.Video, 0)
	for _, m := range ms {
		if n := Video(m); n != nil {
			videos = append(videos, n)
		}
	}
	return videos
}
