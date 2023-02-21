package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
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
func BenchmarkUserLoginService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.UserLoginRequest{
			Username: "ljz",
			Password: "20020210",
		}
		userservice.UserLoginService(context.Background(), &req)
	}
}
func BenchmarkUserLoginServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.UserLoginRequest{
				Username: "ljz",
				Password: "20020210",
			}
			userservice.UserLoginService(context.Background(), &req)
		}
	})
}
