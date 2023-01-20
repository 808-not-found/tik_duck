package pack

import (
	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func Video(m *db.Video) *user.Video {
	if m == nil {
		return nil
	}

	return &user.Video{
		Id:            m.ID,
		Author:        &m.Author,
		PlayPath:      m.PlayPath,
		CoverPath:     m.CoverPath,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    m.IsFavorite,
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
