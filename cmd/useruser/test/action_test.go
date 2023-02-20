package uuservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	uuservice "github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	//"github.com/808-not-found/tik_duck/pkg/consts"
	//"github.com/808-not-found/tik_duck/cmd/useruser/pack"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct RelationActionRequest {
//	    1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	    2: i64 ToUserId (go.tag = 'json:"to_user_id"')//对方用户id
//	    3: i32 ActionType (go.tag = 'json:"action_type"')// 1-关注，2-取消关注
//	}
//
//	struct RelationActionResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值-失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	}
func TestUserRelationActionService(t *testing.T) {
	// 构建通用信息

	// 正确情况测试_关注
	PatchConvey("TestUserRelationActionService_normal", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.FollowAction).Return(nil).Build()
		Mock(db.UnFollowAction).Return(nil).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 1,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 正确情况测试_取消关注
	PatchConvey("TestUserRelationActionService_normal", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.FollowAction).Return(nil).Build()
		Mock(db.UnFollowAction).Return(nil).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 2,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 未登陆
	PatchConvey("TestUserRelationActionService_WithoutToken", t, func() {
		//设置期待值
		expectstatusCode := int32(3002)
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 0}, nil).Build()
		Mock(db.FollowAction).Return(nil).Build()
		Mock(db.UnFollowAction).Return(nil).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 1,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 验证Token失败
	PatchConvey("TestUserRelationActionService_WrongToken", t, func() {
		//设置期待值
		expectstatusCode := int32(3001)
		expectErr := allerrors.ErrTestnotnil()
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 0}, allerrors.ErrTestnotnil()).Build()
		Mock(db.FollowAction).Return(nil).Build()
		Mock(db.UnFollowAction).Return(nil).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 1,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectErr, err)
	})

	//FollowAction失败
	PatchConvey("TestUserRelationActionService_WrongFollowAction", t, func() {
		//设置期待值
		expectstatusCode := int32(3003)
		expectErr := allerrors.ErrTestnotnil()
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.FollowAction).Return(expectErr).Build()
		Mock(db.UnFollowAction).Return(nil).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 1,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectErr, err)
	})

	// UnFollowAction失败
	PatchConvey("TestUserRelationActionService_WrongUnFollowAction", t, func() {
		//设置期待值
		expectstatusCode := int32(3004)
		expectErr := allerrors.ErrTestnotnil()
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.FollowAction).Return(nil).Build()
		Mock(db.UnFollowAction).Return(expectErr).Build()

		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "123123",
			ToUserId:   1,
			ActionType: 2,
		}

		//调用函数
		res, err := uuservice.UserRelationActionService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectErr, err)
	})
}
