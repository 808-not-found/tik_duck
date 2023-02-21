package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	"github.com/808-not-found/tik_duck/pkg/salt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

//	struct UserLoginRequest {
//	    1: string Username (go.tag = 'json:"username"') //登录用户名
//	    2: string Password (go.tag = 'json:"password"') //登录密码
//	}
//
//	struct UserLoginResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	    3: i64 UserId (go.tag = 'json:"user_id"') //用户id
//	    4: string Token (go.tag = 'json:"token"') //用户鉴权token
//	}

func TestUserLoginService(t *testing.T) {
	//构建通用信息
	nowTime := time.Now()
	retUser := db.User{
		ID:            1,
		CreateTime:    nowTime,
		Name:          "蒂萨久",
		FollowCount:   0,
		FollowerCount: 0,
		Password:      "114514",
		Salt:          "1919810",
	}

	//正常情况测试
	PatchConvey("TestUserLoginService_normal", t, func() {
		//设置期待值
		expectID := int64(1)
		expectToken := "right_Token"
		var expectMsg string = consts.Success
		expectstatusCode := int32(0)

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.QueryUser).Return(&retUser, nil).Build()
		Mock(salt.PasswordsMatch).Return(true).Build()
		Mock(jwt.GenToken).Return(expectToken, nil).Build()

		//设置传入参数
		req := user.UserLoginRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserLoginService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, nil, err)
	})

	//用户密码错误
	PatchConvey("TestUserLoginService_WrongPassword", t, func() {
		//设置期待值
		expectID := int64(0)
		expectToken := ""
		expectstatusCode := int32(1006)
		var expectMsg string

		// 设定mock函数
		Mock(salt.PasswordsMatch).Return(false).Build()
		Mock(db.QueryUser).Return(&retUser, nil).Build()
		Mock(jwt.GenToken).Return(expectToken, nil).Build()
		//设置传入参数
		req := user.UserLoginRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserLoginService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, nil, err)
	})

	//用户不存在
	PatchConvey("TestUserLoginService_WrongUsername", t, func() {
		//设置期待值
		expectID := int64(0)
		expectToken := ""
		expectstatusCode := int32(1005)
		expectErr := gorm.ErrRecordNotFound
		var expectMsg string

		// 设定mock函数
		Mock(salt.PasswordsMatch).Return(false).Build()
		Mock(db.QueryUser).Return(&retUser, gorm.ErrRecordNotFound).Build()
		Mock(jwt.GenToken).Return(expectToken, nil).Build()
		//设置传入参数
		req := user.UserLoginRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserLoginService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, expectErr, err)
	})

	//Token返回错误
	PatchConvey("TestUserLoginService_WrongToken", t, func() {
		//设置期待值
		expectID := int64(0)
		expectToken := ""
		expectstatusCode := int32(1007)
		expecterr := allerrors.ErrTestnotnil()
		var expectMsg string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(salt.PasswordsMatch).Return(true).Build()
		Mock(db.QueryUser).Return(&retUser, nil).Build()
		Mock(jwt.GenToken).Return("", allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserLoginRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserLoginService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, expecterr, err)
	})

	//用户名登陆时的未知错误
	PatchConvey("TestUserLoginService_UnknownWrong", t, func() {
		//设置期待值
		expectID := int64(0)
		expectToken := ""
		expectstatusCode := int32(1004)
		expecterr := allerrors.ErrTestnotnil()
		var expectMsg string
		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.QueryUser).Return(&db.User{}, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserLoginRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserLoginService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, expecterr, err)
	})
}
