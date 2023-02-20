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
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

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
		expectUserInfo := user.User{
			Id:            2,
			FollowCount:   nil,
			FollowerCount: nil,
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
		resStatusCode, _, resUserInfo, err := userservice.UserInfoService(context.Background(), &req)

		//对比返回值
		assert.Equal(t, expectstatusCode, resStatusCode)
		assert.Equal(t, expectUserInfo.Id, resUserInfo.Id)
		assert.Equal(t, expectUserInfo.FollowCount, resUserInfo.FollowCount)
		assert.Equal(t, expectUserInfo.FollowerCount, resUserInfo.FollowerCount)
		assert.Equal(t, err, nil)
	})
}
