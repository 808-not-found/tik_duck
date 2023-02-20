package userplatservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

// struct FavoriteActionRequest {
//     1: string Token (go.tag = 'json:"token"')
//     2: i64 VideoId (go.tag = 'json:"video_id"')
//     3: i32 ActionType (go.tag = 'json:"action_type"')
// }

//	struct FavoriteActionResponse {
//	    1: i32 StatusCode   (go.tag = 'json:"status_code"')        //状态码，0-成功，其他值失败
//	    2: string StatusMsg (go.tag = 'json:"status_msg"')        // 返回状态描述
//	}
func TestUserFavoriteActionService(t *testing.T) {
	//鉴权失败
	PatchConvey("TestMockUserplatFavoriteAction_WorryClaim", t, func() {
		expectstatusCode := int32(2040)
		expectStatusMsg := ""
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteActionRequest{
			Token:      "3",
			VideoId:    999,
			ActionType: 1,
		}
		res, err := userplatservice.UserFavoriteActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, err, nil)
	})
	//点赞失败
	PatchConvey("TestMockUserplatFavoriteAction_Worrylike", t, func() {
		expectstatusCode := int32(2041)
		expectStatusMsg := ""
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.LikeAction).Return(allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteActionRequest{
			Token:      "3",
			VideoId:    999,
			ActionType: 1,
		}
		res, err := userplatservice.UserFavoriteActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//取消点赞失败
	PatchConvey("TestMockUserplatFavoriteAction_Worryunlike", t, func() {
		expectstatusCode := int32(2041)
		expectStatusMsg := ""
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UnLikeAction).Return(allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteActionRequest{
			Token:      "3",
			VideoId:    999,
			ActionType: 2,
		}
		res, err := userplatservice.UserFavoriteActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//点赞成功
	PatchConvey("TestMockUserplatFavoriteAction_Normal", t, func() {
		expectstatusCode := int32(0)
		expectStatusMsg := ""
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.LikeAction).Return(nil).Build()
		req := userplat.FavoriteActionRequest{
			Token:      "3",
			VideoId:    999,
			ActionType: 1,
		}
		res, err := userplatservice.UserFavoriteActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, err, nil)
	})
}
