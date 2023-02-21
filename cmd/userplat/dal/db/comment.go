package db

import (
	"context"
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

func CommentAction(ctx context.Context, myID int64, vdID int64, commentText string) (*Comment, error) {
	// 增加commentcount
	// var myUser *User
	var res *Comment // 返回值为一条评论内容
	//var commentlist *Comment
	var myVideo *Video
	conn := DB.WithContext(ctx).Where("id = ?", vdID).First(&myVideo).Update("comment_count", myVideo.CommentCount+1)
	if err := conn.Error; err != nil {
		return res, err
	}
	// 增加一条记录到Comment表
	comment := Comment{
		UserID:  myID,
		VideoID: vdID,
		Content: commentText,
	}
	conn = DB.WithContext(ctx).Create(&comment)
	//res = &comment
	if err := conn.Error; err != nil {
		return res, err
	}
	// // 找到所有和 myID 相关的记录
	// conn := DB.WithContext(ctx).Where("user_id = ?", myID).Find(&commentlist)
	// if err := conn.Error; err != nil {
	// 	return res, err
	// }
	// 找到对应的视频结构体
	conn = DB.WithContext(ctx).Where("id = ?", myID).Find(&res)
	if err := conn.Error; err != nil {
		return res, err
	}
	return res, nil
}

func UnCommentAction(ctx context.Context, myID int64, vdID int64, commentID int64) error {
	// 减少commentcount
	// var myUser *User
	var myVideo *Video
	conn := DB.WithContext(ctx).Where("id = ?", vdID).First(&myVideo).Update("comment_count", myVideo.CommentCount-1)
	if err := conn.Error; err != nil {
		return err
	}
	// 减少一条记录到Comment表
	comment := Comment{
		UserID:  myID,
		VideoID: vdID,
		ID:      commentID, // 应该是删除的commentID的内容的ID
	}
	conn = DB.WithContext(ctx).Delete(&comment)
	if err := conn.Error; err != nil {
		return err
	}
	return nil
}

func GetCommentList(ctx context.Context, myID int64, vdID int64) ([]*Comment, error) {
	var res []*Comment
	// 找到所有和 myID 相关的记录
	var commentList []*Comment
	conn := DB.WithContext(ctx).Where("user_id = ?", myID).Find(&commentList)
	if err := conn.Error; err != nil {
		return res, err
	}

	// 获取所有的评论 ID
	setList := make(map[int64]bool)
	for _, value := range commentList {
		if value != nil {
			setList[value.VideoID] = true // nolint:all
		}
	}
	var commentIDList []int64
	for k := range setList {
		commentIDList = append(commentIDList, k)
	}

	// 找到所有对应的视频结构体
	conn = DB.WithContext(ctx).Where("id = ?", commentIDList).Find(&res)
	if err := conn.Error; err != nil {
		return res, err
	}
	return res, nil
}
