package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

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

func UserGetFeed(ctx context.Context, latestTime time.Time) ([]*Video, int64, error) {
	// 初始化数据
	if latestTime.IsZero() {
		latestTime = time.Now()
	}
	const limit = 30

	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Limit(limit).Where("publish_time <= ?", latestTime).Find(&videoList)
	if err := conn.Error; err != nil {
		nilTime := time.Time{}
		return nil, nilTime.Unix(), err
	}

	// 获取最早时间
	var firstVideo Video
	conn.Order("publish_time").First(&firstVideo)
	nextTime := firstVideo.PublishTime.Unix()

	return videoList, nextTime, nil
}

func UserPublishList(ctx context.Context, userID int32) ([]*Video, error) {
	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("author_id = ?", userID).Find(&videoList)
	if err := conn.Error; err != nil {
		return nil, err
	}

	// 返回
	return videoList, nil
}

func UserPublishAction(ctx context.Context, userID int64, filePath string, coverPath string, title string) error {
	video := Video{AuthorID: userID, FilePath: filePath, CoverPath: coverPath, Title: title}
	res := DB.Create(&video)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
