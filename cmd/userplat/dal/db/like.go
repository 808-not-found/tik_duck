package db

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/pkg/consts"
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LikeTime time.Time `gorm:"column:like_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;default:0;NOT NULL"`
	VideoID  int64     `gorm:"column:video_id;default:0;NOT NULL"`
}

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

type User struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	Name          string    `gorm:"column:name;unique;NOT NULL"`
	Password      string    `gorm:"column:password;NOT NULL"`
	Salt          string    `gorm:"column:salt;NOT NULL"`
	FollowCount   int64     `gorm:"column:follow_count;default:0;NOT NULL"`
	FollowerCount int64     `gorm:"column:follower_count;default:0;NOT NULL"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func (like *Like) TableName() string {
	return consts.LikeTableName
}

// 传入用户id 返回用户信息.
func GetUser(ctx context.Context, userID int64) (User, error) {
	res := User{}
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// myID 点赞 vdID.
func LikeAction(ctx context.Context, myID int64, vdID int64) error {
	// 增加favoritecount
	// var myUser *User
	var myVideo *Video
	conn := DB.WithContext(ctx).Where("id = ?", vdID).First(&myVideo).Update("id", myVideo.FavoriteCount+1)
	if err := conn.Error; err != nil {
		return err
	}
	// 增加一条记录到Like表
	like := Like{
		UserID:  myID,
		VideoID: vdID,
	}
	conn = DB.WithContext(ctx).Create(like)
	if err := conn.Error; err != nil {
		return err
	}
	return nil
}

func UnLikeAction(ctx context.Context, myID int64, vdID int64) error {
	// 减少favoritecount
	// var myUser *User
	var myVideo *Video
	conn := DB.WithContext(ctx).Where("id = ?", vdID).First(&myVideo).Update("id", myVideo.FavoriteCount-1)
	if err := conn.Error; err != nil {
		return err
	}
	// 删除Like表中的一条记录
	like := Like{
		UserID:  myID,
		VideoID: vdID,
	}
	conn = DB.WithContext(ctx).Create(like)
	if err := conn.Error; err != nil {
		return err
	}
	return nil
}

func GetFavoriteList(ctx context.Context, userID int64) ([]*Video, error) {
	var res []*Video
	// 找到所有和 userID 相关的记录
	var favoriteList []*Like
	conn := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&favoriteList)
	if err := conn.Error; err != nil {
		return res, err
	}

	// 获取所有的视频 ID
	var favoriteIDList []int64
	for _, value := range favoriteList {
		favoriteIDList = append(favoriteIDList, value.VideoID)
	}

	// 找到所有对应的视频结构体
	conn = DB.WithContext(ctx).Where("id = ?", favoriteIDList).Find(&res)
	if err := conn.Error; err != nil {
		return res, err
	}

	return res, nil
}
