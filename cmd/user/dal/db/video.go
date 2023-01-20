package db

import (
	"context"
	"net/http"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID   int64 `json:"id"`
	User User  `json:"author"`
	// PlayPath      string `json:play_path`
	// CoverPath     string `json:cover_path`
	// FavoriteCount int64  `json:favorite_count`
	// CommentCount  int64  `json:comment_count`
	// IsFavorite    bool   `json:is_favorite`
	// Title         string `json:title`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func UserGetFeed(ctx context.Context, latestTime int64, token string) (int32, string, []*Video, int64) {
	var statusCode int32
	var statusMsg string
	var videoList []*Video
	var nextTime int64
	const num = 30

	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	conn := DB.WithContext(ctx).Model(&Video{}).Limit(num).Where("publish_time <= ?", latestTime)

	statusCode = http.StatusAccepted
	statusMsg = "success"
	conn.Pluck("file_path", &videoList)
	conn.Order("publish_time").First(&Video{}).Pluck("publish_time", &nextTime)

	return statusCode, statusMsg, videoList, nextTime
}
