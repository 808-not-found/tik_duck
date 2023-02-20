package userservice_test

import (
	"context"
	//"gorm.io/gorm"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct UserRequest {
//	    1: i64 UserId (go.tag = 'json:"user_id"') //用户id
//	    2: string Token (go.tag = 'json:"token"') //用户鉴权token
//	}
//
//	struct UserResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值-失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	    3: User User (go.tag = 'json:"user"') //用户信息
//	}
func TestUserInfoService(t *testing.T) {
	//构建通用信息
	nowTime := time.Now()
	retUser1 := db.User{
		ID: 1,

		CreateTime:    nowTime,
		Name:          "蒂萨久",
		FollowCount:   0,
		FollowerCount: 0,
		Password:      "114514",
		Salt:          "1919810",
	}
	retUser2 := db.User{
		ID:            2,
		CreateTime:    nowTime,
		Name:          "皮卡皮",
		FollowCount:   0,
		FollowerCount: 0,
		Password:      "808808",
		Salt:          "2002808",
	}

	//正常情况测试
	PatchConvey("TestUserInfoService_normal", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
		expectstatusMsg := consts.Success
		expectUserInfo := user.User{
			Id:            2,
			Name:          "皮卡皮",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}

		// 设定mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1, Username: "皮卡皮"}, nil).Build()
		Mock(db.QueryUser).Return(&retUser1, nil).Build()
		Mock(db.GetUser).Return(retUser2, nil).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{
			Id:            2,
			Name:          "皮卡皮",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}, nil).Build()

		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}

		//调用函数
		resStatusCode, resStatusMsg, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectstatusMsg, resStatusMsg)
		assert.Equal(t, &expectUserInfo, resUserInfo)
		assert.Equal(t, nil, err)
	})

	//用户鉴权失败
	PatchConvey("TestUserInfoService_WrongToken", t, func() {
		//设置期待值
		expectstatusCode := int32(1008)
		expectstatusMsg := ""
		expectErr := allerrors.ErrTestnotnil()
		var expectUserInfo *user.User

		// 设定mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}

		//调用函数
		resStatusCode, resStatusMsg, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectstatusMsg, resStatusMsg)
		assert.Equal(t, expectUserInfo, resUserInfo)
		assert.Equal(t, expectErr, err)
	})

	//自己id查询失败
	PatchConvey("TestUserInfoService_WrongFrom", t, func() {
		//设置期待值
		expectstatusCode := int32(1009)
		expectstatusMsg := ""
		expectErr := allerrors.ErrTestnotnil()
		var expectUserInfo *user.User

		// 设定mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1, Username: "皮卡皮"}, nil).Build()
		Mock(db.QueryUser).Return(&retUser1, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}

		//调用函数
		resStatusCode, resStatusMsg, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectstatusMsg, resStatusMsg)
		assert.Equal(t, expectUserInfo, resUserInfo)
		assert.Equal(t, expectErr, err)
	})

	//查询对方id失败
	PatchConvey("TestUserInfoService_WrongAim", t, func() {
		//设置期待值
		expectstatusCode := int32(1010)
		expectstatusMsg := ""
		expectErr := allerrors.ErrTestnotnil()
		var expectUserInfo *user.User

		// 设定mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1, Username: "皮卡皮"}, nil).Build()
		Mock(db.QueryUser).Return(&retUser1, nil).Build()
		Mock(db.GetUser).Return(retUser2, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}

		//调用函数
		resStatusCode, resStatusMsg, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectstatusMsg, resStatusMsg)
		assert.Equal(t, expectUserInfo, resUserInfo)
		assert.Equal(t, expectErr, err)
	})

	//查询两人关系失败
	PatchConvey("TestUserInfoService_WrongRelation", t, func() {
		//设置期待值
		expectstatusCode := int32(1011)
		expectstatusMsg := ""
		expectErr := allerrors.ErrTestnotnil()
		var expectUserInfo *user.User

		// 设定mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1, Username: "皮卡皮"}, nil).Build()
		Mock(db.QueryUser).Return(&retUser1, nil).Build()
		Mock(db.GetUser).Return(retUser2, nil).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{}, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}

		//调用函数
		resStatusCode, resStatusMsg, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectstatusMsg, resStatusMsg)
		assert.Equal(t, expectUserInfo, resUserInfo)
		assert.Equal(t, expectErr, err)
	})
}
