package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	MessageTime time.Time `gorm:"column:message_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID  int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID    int64     `gorm:"column:to_user_id;NOT NULL"`
	Content     string    `gorm:"column:content;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func GetMsg(ctx context.Context, myID int64, toID int64, t *int64) (res []*Message, err error) {
	curTime := time.Unix(*t, 0)
	conn := DB.WithContext(ctx).Where("created_at <= ?", curTime).Find(&res)
	if err = conn.Error; err != nil {
		return res, err
	}
	return res, nil
}

func PostMsg(ctx context.Context, myID int64, toID int64, content *string) (err error) {
	message := Message{
		FromUserID: myID,
		ToUserID: toID,
		Content: *content,
	}
	conn := DB.WithContext(ctx).Create(&message)
	if err = conn.Error; err != nil {
		return err
	}

	return nil
}
