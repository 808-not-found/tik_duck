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

type Like struct {
	gorm.Model
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LikeTime time.Time `gorm:"column:like_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;default:0;NOT NULL"`
	VideoID  int64     `gorm:"column:video_id;default:0;NOT NULL"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func (v *Like) TableName() string {
	return consts.LikeTableName
}

func UserGetFeed(ctx context.Context, latestTime *int64) ([]*Video, error) {
	// 初始化数据
	var curTime time.Time
	// 客户端可能会传一个很大的时间来表示首次请求

	if latestTime == nil || *latestTime == 0 || *latestTime > 1076997151921 {
		curTime = time.Now()
	} else {
		curTime = time.Unix(*latestTime, 0)
	}

	curTime = curTime.Add(8 * time.Hour)
	const limit = 30

	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("publish_time <= ?", curTime).
		Limit(limit).Order("publish_time desc").Find(&videoList)
	if err := conn.Error; err != nil {
		return nil, err
	}

	return videoList, nil
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

func IsFavorite(ctx context.Context, userID int64, vdeioID int64) error {
	var like *Like
	conn := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, vdeioID).First(&like)
	if err := conn.Error; err != nil {
		return err
	}

	return nil
}
