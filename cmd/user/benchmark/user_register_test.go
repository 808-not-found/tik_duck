package userservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

func TestUserRegisterService(t *testing.T) {
	//构建通用信息

	//正常情况测试
	PatchConvey("TestUserRegisterService_normal", t, func() {
		//设置期待值
		expectID := int64(1)
		expectToken := "right_Token"
		var expectMsg string = consts.Success
		expectstatusCode := int32(0)

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.CheckUserExist).Return(false, nil).Build()
		Mock(db.CreateUser).Return(nil).Build()
		Mock(jwt.GenToken).Return(expectToken, nil).Build()
		Mock(db.GetUserID).Return(int64(1), nil).Build()
		//设置传入参数
		req := user.UserRegisterRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserRegisterService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, nil, err)
	})

	//数据库创建失败
	PatchConvey("TestUserRegisterService_WrongCreate", t, func() {
		expectID := int64(0)
		expectToken := ""
		var expectMsg string
		expectstatusCode := int32(1002)
		expectErr := allerrors.ErrTestnotnil()

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.CheckUserExist).Return(false, nil).Build()
		Mock(db.CreateUser).Return(allerrors.ErrTestnotnil()).Build()
		Mock(jwt.GenToken).Return(expectToken, nil).Build()
		Mock(db.GetUserID).Return(int64(1), nil).Build()
		//设置传入参数
		req := user.UserRegisterRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserRegisterService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, expectErr, err)
	})

	//Token生成失败
	PatchConvey("TestUserRegisterService_WrongToken", t, func() {
		expectID := int64(0)
		expectToken := ""
		var expectMsg string
		expectstatusCode := int32(1003)
		expectErr := allerrors.ErrTestnotnil()

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.CheckUserExist).Return(false, nil).Build()
		Mock(db.CreateUser).Return(nil).Build()
		Mock(jwt.GenToken).Return("", allerrors.ErrTestnotnil()).Build()
		Mock(db.GetUserID).Return(int64(1), nil).Build()
		//设置传入参数
		req := user.UserRegisterRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserRegisterService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, expectErr, err)
	})

	//注册重复
	PatchConvey("TestUserRegisterService_WrongToken", t, func() {
		expectID := int64(0)
		expectToken := ""
		var expectMsg string
		expectstatusCode := int32(1012)

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.CheckUserExist).Return(true, nil).Build()
		Mock(db.CreateUser).Return(nil).Build()
		Mock(jwt.GenToken).Return("", allerrors.ErrTestnotnil()).Build()
		Mock(db.GetUserID).Return(int64(1), nil).Build()
		//设置传入参数
		req := user.UserRegisterRequest{
			Username: "蒂萨久",
			Password: "114514",
		}

		//调用函数
		statusCode, resStatusMsg, resID, resToken, err := userservice.UserRegisterService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectID, resID)
		assert.Equal(t, expectToken, resToken)
		assert.Equal(t, expectMsg, resStatusMsg)
		assert.Equal(t, expectstatusCode, statusCode)
		assert.Equal(t, nil, err)
	})
}
