package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID            int64     `json:"id"`
	Author        user.User `json:"author"`
	PlayPath      string
	CoverPath     string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
	PublishTime   int64 `json:"publish_time"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func UserGetFeed(ctx context.Context, latestTime int64, token string) ([]*Video, int64, error) {
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	const limit = 30

	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Limit(limit).Where("publish_time <= ?", latestTime).Find(&videoList)
	if err := conn.Error; err != nil {
		return nil, 0, err
	}

	// 获取当前列表的最早的视频
	var firstVideo Video
	conn.Order("publish_time").First(&firstVideo)
	return videoList, firstVideo.PublishTime, nil
}
