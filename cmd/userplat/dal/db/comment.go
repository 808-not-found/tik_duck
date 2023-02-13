package db

import (
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CommentTime time.Time `gorm:"column:comment_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID      int64     `gorm:"column:user_id;default:0;NOT NULL"`
	VideoID     int64     `gorm:"column:video_id;default:0;NOT NULL"`
	Content     string    `gorm:"column:content;NOT NULL"`
}

func (comment *Comment) TableName() string {
	return consts.CommentTableName
}

func CommentAction(myID int64, vdID int64, commentText *string) error {
	return nil
}

func UnCommentAction(myID int64, vdID int64, commentID int64) error {
	return nil
}
