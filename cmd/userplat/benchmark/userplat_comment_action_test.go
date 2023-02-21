package userplatservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct CommentActionRequest {
//	   1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	   2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
//	   3: i32 ActionType (go.tag = 'json:"action_type"')// 1- 发布评论，2- 删除评论
//	   4: optional string CommentText (go.tag = 'json:"comment_text"')//用户填写的评论内容，在action_type=1 的时候使用
//	   5: optional i64 CommentId (go.tag = 'json:"comment_id"')//要删除的评论id,在action_type=2的时候使用
//	}
//
//	struct CommentActionResponse {
//	   1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
//	   2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	   3: optional Comment Comment (go.tag = 'json:"comment"')//评论成功返回评论内容，不需要重新拉取整个列表
//	}
func TestUserCommentActionService(t *testing.T) {
	rettext := "I hate you"
	// 用户鉴权失败
	PatchConvey("TestMockUserplatCommentAction_WorryClaim", t, func() {
		expectstatusCode := int32(2034)
		var expectStatusMsg *string
		var expectComment *userplat.Comment
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentActionRequest{
			Token:       "3",
			VideoId:     999,
			ActionType:  1,
			CommentText: &rettext,
		}
		res, err := userplatservice.UserCommentActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.Comment)
		assert.Equal(t, err, nil)
	})
	// 用户评论数据库失败
	PatchConvey("TestMockUserplatCommentAction_WorrydbComment", t, func() {
		expectstatusCode := int32(2035)
		var expectStatusMsg *string
		var expectComment *userplat.Comment
		nowTime := time.Now()
		retcomment := db.Comment{
			ID:          999,
			CommentTime: nowTime,
			UserID:      222,
			VideoID:     777,
			Content:     "I love you",
		}
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.CommentAction).Return(&retcomment, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentActionRequest{
			Token:       "2333",
			VideoId:     999,
			ActionType:  1,
			CommentText: &rettext,
		}
		res, err := userplatservice.UserCommentActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.Comment)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	// 用户删除评论数据库失败
	PatchConvey("TestMockUserplatCommentAction_WorrydbUnComment", t, func() {
		expectstatusCode := int32(2035)
		reqCommentID := int64(1)
		var expectStatusMsg *string
		var expectComment *userplat.Comment
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UnCommentAction).Return(allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentActionRequest{
			Token:       "2333",
			VideoId:     999,
			ActionType:  2,
			CommentId:   &reqCommentID,
			CommentText: &rettext,
		}
		res, err := userplatservice.UserCommentActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.Comment)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	// 用户评论封装失败
	PatchConvey("TestMockUserplatCommentAction_WorryrpcComment", t, func() {
		expectstatusCode := int32(2036)
		var expectStatusMsg *string
		var expectComment *userplat.Comment
		nowTime := time.Now()
		retcomment := db.Comment{
			ID:          999,
			CommentTime: nowTime,
			UserID:      222,
			VideoID:     777,
			Content:     "I love you",
		}
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.CommentAction).Return(&retcomment, nil).Build()
		Mock(pack.Comment).Return(expectComment, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentActionRequest{
			Token:       "2333",
			VideoId:     999,
			ActionType:  1,
			CommentText: &rettext,
		}
		res, err := userplatservice.UserCommentActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.Comment)
		assert.Equal(t, allerrors.ErrTestnotnil(), err)
	})
	// 用户评论成功
	PatchConvey("TestMockUserplatCommentAction_normal", t, func() {
		expectstatusCode := int32(0)
		var expectStatusMsg *string
		nowTime := time.Now()
		retcomment := db.Comment{
			ID:          999,
			CommentTime: nowTime,
			UserID:      222,
			VideoID:     777,
			Content:     "I love you",
		}
		retUser := userplat.User{
			Id:            2,
			Name:          "Hentai",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}
		expectcomment := &userplat.Comment{
			Id:         9,
			User:       &retUser,
			Content:    rettext,
			CreateDate: "2022.2.22",
		}
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.CommentAction).Return(&retcomment, nil).Build()
		Mock(pack.Comment).Return(expectcomment, nil).Build()
		req := userplat.CommentActionRequest{
			Token:       "2333",
			VideoId:     999,
			ActionType:  1,
			CommentText: &rettext,
		}
		res, err := userplatservice.UserCommentActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectcomment, res.Comment)
		assert.Equal(t, err, nil)
	})
}
