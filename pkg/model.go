package model

import (
	"time"
)

type Video struct {
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AuthorID      int64     `gorm:"column:author_id;NOT NULL"`
	PublishTime   time.Time `gorm:"column:publish_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FilePath      string    `gorm:"column:file_path;NOT NULL"`
	CoverPath     string    `gorm:"column:cover_path;NOT NULL"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	Title         string    `gorm:"column:title;NOT NULL"`
}

func (m *Video) TableName() string {
	return "video"
}

type User struct {
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	Name          string    `gorm:"column:name;unique;NOT NULL"`
	Password      string    `gorm:"column:password;NOT NULL"`
	Salt          string    `gorm:"column:salt;NOT NULL"`
	FollowCount   int64     `gorm:"column:follow_count;default:0;NOT NULL"`
	FollowerCount int64     `gorm:"column:follower_count;default:0;NOT NULL"`
}

func (m *User) TableName() string {
	return "user"
}

type Follow struct {
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
}

func (m *Follow) TableName() string {
	return "follow"
}

type Comment struct {
	ID          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CommentTime time.Time `gorm:"column:comment_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID      int64     `gorm:"column:user_id;NOT NULL"`
	VideoID     int64     `gorm:"column:video_id;NOT NULL"`
	Content     string    `gorm:"column:content;NOT NULL"`
}

func (m *Comment) TableName() string {
	return "comment"
}

type Message struct {
	ID          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	MessageTime time.Time `gorm:"column:message_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID  int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID    int64     `gorm:"column:to_user_id;NOT NULL"`
	Content     string    `gorm:"column:content;NOT NULL"`
}

func (m *Message) TableName() string {
	return "message"
}

type Like struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LikeTime time.Time `gorm:"column:like_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;NOT NULL"`
	VideoID  int64     `gorm:"column:video_id;NOT NULL"`
}

func (m *Like) TableName() string {
	return "like"
}
