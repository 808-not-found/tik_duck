package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

//	type Video struct {
//			gorm.Model
//			ID            int64     `json:"id"`
//			Author        user.User `json:"author"`
//			PlayPath      string
//			CoverPath     string
//			FavoriteCount int64
//			CommentCount  int64
//			IsFavorite    bool
//			Title         string
//			PublishTime   int64 `json:"publish_time"`
//		}
type Video struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AuthorID      int64     `gorm:"column:author_id;NOT NULL"`
	PublishTime   time.Time `gorm:"column:publish_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FilePath      string    `gorm:"column:file_path;NOT NULL"`
	CoverPath     string    `gorm:"column:cover_path;NOT NULL"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	Title         string    `gorm:"column:title;NOT NULL"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// 更改了video之后 有一点小问题 数据库里面的video中 PublishTime 是 time类型
// 而你原来的是int64类型 我先改成了time类型 先过审 你回头看看是要哪个.
func UserGetFeed(ctx context.Context, latestTime time.Time, token string) ([]*Video, time.Time, error) {
	if latestTime.IsZero() {
		latestTime = time.Now()
	}
	const limit = 30

	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Limit(limit).Where("publish_time <= ?", latestTime).Find(&videoList)
	if err := conn.Error; err != nil {
		nilTime := time.Time{}
		return nil, nilTime, err
	}

	// 获取当前列表的最早的视频
	var firstVideo Video
	conn.Order("publish_time").First(&firstVideo)
	return videoList, firstVideo.PublishTime, nil
}

// 原版
// func UserGetFeed(ctx context.Context, latestTime int64, token string) ([]*Video, int64, error) {
// 	if latestTime == 0 {
// 		latestTime = time.Now().Unix()
// 	}
// 	const limit = 30

// 	// 获取视频列表
// 	var videoList []*Video
// 	conn := DB.WithContext(ctx).Model(&Video{}).Limit(limit).Where("publish_time <= ?", latestTime).Find(&videoList)
// 	if err := conn.Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// 获取当前列表的最早的视频
// 	var firstVideo Video
// 	conn.Order("publish_time").First(&firstVideo)
// 	return videoList, firstVideo.PublishTime, nil
// }
