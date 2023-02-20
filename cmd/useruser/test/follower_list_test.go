package uuservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	//"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/cmd/useruser/pack"
	uuservice "github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

func TestUserRelationFollowerList(t *testing.T) {
	// 构建通用信息
	nowTime := time.Now()
	retUserList := make([]*db.User, 0)
	retUser1 := db.User{
		ID:            1,
		CreateTime:    nowTime,
		Name:          "蒂萨久",
		FollowCount:   0,
		FollowerCount: 0,
		Password:      "114514",
		Salt:          "1919810",
	}
	retUserList = append(retUserList, &retUser1)
	// 正确情况测试_关注
	PatchConvey("TestUserRelationFollowerList_normal", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
		expectUserList := make([]*useruser.User, 0)
		expectUserList = append(expectUserList, &useruser.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		},
		)
		var expectMsg *string
		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.GetFollowerList).Return(retUserList, nil).Build()
		Mock(pack.Users).Return(expectUserList, nil).Build()

		//设置传入参数
		req := useruser.RelationFollowerListRequest{
			UserId: 1,
			Token:  "xxxyyyxxx",
		}

		//调用函数
		res, err := uuservice.UserRelationFollowerList(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectUserList, res.UserList)
		assert.Equal(t, nil, err)
	})

	// 验证Token失败
	PatchConvey("TestUserRelationFollowerList_WrongToken", t, func() {
		//设置期待值
		expectstatusCode := int32(3009)
		expectErr := allerrors.ErrTestnotnil()
		var expectUserList []*useruser.User
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, expectErr).Build()
		Mock(db.GetFollowerList).Return(retUserList, nil).Build()
		Mock(pack.Users).Return(expectUserList, nil).Build()

		//设置传入参数
		req := useruser.RelationFollowerListRequest{
			UserId: 1,
			Token:  "xxxyyyxxx",
		}

		//调用函数
		res, err := uuservice.UserRelationFollowerList(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectUserList, res.UserList)
		assert.Equal(t, expectErr, err)
	})

	// 未登陆
	PatchConvey("TestUserRelationFollowerList_WithoutToken", t, func() {
		//设置期待值
		expectstatusCode := int32(3010)
		var expectUserList []*useruser.User
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 0}, nil).Build()
		Mock(db.GetFollowerList).Return(retUserList, nil).Build()
		Mock(pack.Users).Return(expectUserList, nil).Build()

		//设置传入参数
		req := useruser.RelationFollowerListRequest{
			UserId: 1,
			Token:  "xxxyyyxxx",
		}

		//调用函数
		res, err := uuservice.UserRelationFollowerList(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectUserList, res.UserList)
		assert.Equal(t, nil, err)
	})

	//查询数据库失败
	PatchConvey("TestUserRelationFollowerList_WrongDB", t, func() {
		//设置期待值
		expectstatusCode := int32(3011)
		expectErr := allerrors.ErrTestnotnil()
		var expectUserList []*useruser.User
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.GetFollowerList).Return(retUserList, expectErr).Build()
		Mock(pack.Users).Return(expectUserList, nil).Build()

		//设置传入参数
		req := useruser.RelationFollowerListRequest{
			UserId: 1,
			Token:  "xxxyyyxxx",
		}

		//调用函数
		res, err := uuservice.UserRelationFollowerList(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectUserList, res.UserList)
		assert.Equal(t, expectErr, err)
	})

	// 数据封装失败
	PatchConvey("TestUserRelationFollowerList_WrongPack", t, func() {
		//设置期待值
		expectstatusCode := int32(3012)
		expectErr := allerrors.ErrTestnotnil()
		var expectUserList []*useruser.User
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.GetFollowerList).Return(retUserList, nil).Build()
		Mock(pack.Users).Return(expectUserList, expectErr).Build()

		//设置传入参数
		req := useruser.RelationFollowerListRequest{
			UserId: 1,
			Token:  "xxxyyyxxx",
		}

		//调用函数
		res, err := uuservice.UserRelationFollowerList(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectUserList, res.UserList)
		assert.Equal(t, expectErr, err)
	})
}
