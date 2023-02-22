package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
)

func BenchmarkUserRegisterService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.UserRegisterRequest{
			Username: "蒂萨久",
			Password: "114514",
		}
		_, _, _, _, _ = userservice.UserRegisterService(context.Background(), &req)
	}
}
func BenchmarkUserRegisterServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.UserRegisterRequest{
				Username: "蒂萨久",
				Password: "114514",
			}
			_, _, _, _, _ = userservice.UserRegisterService(context.Background(), &req)
		}
	})
}
