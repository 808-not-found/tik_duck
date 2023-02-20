package userplatservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct CommentListRequest {
//	    1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	    2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
//	}
//
//	struct CommentListResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	    3: list<Comment> CommentList (go.tag = 'json:"comment_list"')//评论列表
//	}
func TestUserCommentListService(t *testing.T) {
	//用户鉴权失败
	PatchConvey("TestMockUserplatCommentList_WorryClaim", t, func() {
		expectstatusCode := int32(2031)
		var expectStatusMsg *string
		var expectComment []*userplat.Comment
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentListRequest{
			Token:   "3",
			VideoId: 777,
		}
		res, err := userplatservice.UserCommentListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.CommentList)
		assert.Equal(t, err, nil)
	})
	//查询数据库错误
	PatchConvey("TestMockUserplatCommentList_Worrydb", t, func() {
		expectstatusCode := int32(2032)
		var expectStatusMsg *string
		var expectComment []*userplat.Comment
		var retdbcomment []*db.Comment
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetCommentList).Return(retdbcomment, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentListRequest{
			Token:   "3",
			VideoId: 777,
		}
		res, err := userplatservice.UserCommentListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectComment, res.CommentList)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//数据封装错误
	PatchConvey("TestMockUserplatCommentList_Worryrpc", t, func() {
		expectstatusCode := int32(2033)
		var expectStatusMsg *string
		var retdbcomment []*db.Comment
		var expectcomment []*userplat.Comment
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetCommentList).Return(retdbcomment, nil).Build()
		Mock(pack.Comments).Return(expectcomment, allerrors.ErrTestnotnil()).Build()
		req := userplat.CommentListRequest{
			Token:   "3",
			VideoId: 777,
		}
		res, err := userplatservice.UserCommentListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectcomment, res.CommentList)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//正常情况
	PatchConvey("TestMockUserplatCommentList_Normal", t, func() {
		expectstatusCode := int32(0)
		var expectStatusMsg *string
		var retdbcomment []*db.Comment
		retUser := userplat.User{
			Id:            2,
			Name:          "Hentai",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}
		expectcomment := make([]*userplat.Comment, 0)
		expectcomment = append(expectcomment, &userplat.Comment{
			Id:         9,
			User:       &retUser,
			Content:    "I hate you",
			CreateDate: "2022.2.22",
		})
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetCommentList).Return(retdbcomment, nil).Build()
		Mock(pack.Comments).Return(expectcomment, nil).Build()
		req := userplat.CommentListRequest{
			Token:   "3",
			VideoId: 777,
		}
		res, err := userplatservice.UserCommentListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectcomment, res.CommentList)
		assert.Equal(t, err, nil)
	})
}
