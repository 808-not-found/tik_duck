package userservice_test

import (
	"context"
	//"gorm.io/gorm"
	"testing"

	// "github.com/808-not-found/tik_duck/cmd/user/dal/db"
	// "github.com/808-not-found/tik_duck/cmd/user/pack"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
	//"github.com/808-not-found/tik_duck/pkg/allerrors"
	// "github.com/808-not-found/tik_duck/pkg/consts"
	// "github.com/808-not-found/tik_duck/pkg/jwt"
	// . "github.com/bytedance/mockey"
	// "github.com/stretchr/testify/assert"
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
func BenchmarkUserInfoService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.UserRequest{
			Token:  "1231312",
			UserId: 1231231,
		}
		userservice.UserInfoService(context.Background(), &req)
	}
}
func BenchmarkUserInfoServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.UserRequest{
				Token:  "1231312",
				UserId: 1231231,
			}
			userservice.UserInfoService(context.Background(), &req)
		}
	})
}
