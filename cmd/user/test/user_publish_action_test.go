package userservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct PublishActionRequest {
//	    1: string Token (go.tag = 'json:"token"') //用户鉴权token
//	    2: string FilePath (go.tag = 'json:"file_path"')  // 视频路径
//	    3: string CoverPath (go.tag = 'json:"cover_path"') // 封面路径
//	    4: string Title (go.tag = 'json:"title"') //视频标题
//	}
//
//	struct PublishActionResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值-失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	}

func TestUserPublishActionService(t *testing.T) {
	//用户鉴权失败
	PatchConvey("TestMockUserPublishAction_WorryClaims", t, func() {
		expectstatusCode := int32(1024)
		expecterr := allerrors.ErrTestnotnil()
		var expextStatusMsg *string
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := user.PublishActionRequest{
			Token:     "1926",
			FilePath:  "public/0.mp4",
			CoverPath: "public/8.jpg",
			Title:     "17",
		}
		res, err := userservice.UserPublishActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expextStatusMsg, res.StatusMsg)
		assert.Equal(t, expecterr, err)
	})
	//数据写入失败
	PatchConvey("TestMockUserPublishAction_WorryDb", t, func() {
		expectstatusCode := int32(1025)
		expecterr := allerrors.ErrTestnotnil()
		var expextStatusMsg *string
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UserPublishAction).Return(allerrors.ErrTestnotnil()).Build()
		req := user.PublishActionRequest{
			Token:     "1926",
			FilePath:  "public/0.mp4",
			CoverPath: "public/8.jpg",
			Title:     "17",
		}
		res, err := userservice.UserPublishActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expextStatusMsg, res.StatusMsg)
		assert.Equal(t, expecterr, err)
	})
	//正常情况
	PatchConvey("TestMockUserPublishAction_normal", t, func() {
		expectstatusCode := int32(0)
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UserPublishAction).Return(nil).Build()
		req := user.PublishActionRequest{
			Token:     "1926",
			FilePath:  "public/0.mp4",
			CoverPath: "public/8.jpg",
			Title:     "17",
		}
		res, err := userservice.UserPublishActionService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, err, nil)
	})
}
