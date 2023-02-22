package uuservice_test

import (
	"context"
	"testing"

	uuservice "github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

//	struct RelationFriendListRequest {
//	    1: i64 UserId (go.tag = 'json:"user_id"')//用户id
//	    2: string Token (go.tag = 'json:"token"')//用户鉴权token
//	}
//
//	struct RelationFriendListResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	    3: list<User> UserList (go.tag = 'json:"user_list"')//用户列表
//	}

func BenchmarkUserRelationFriendListService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := useruser.RelationFriendListRequest{
			Token:  "1231312",
			UserId: 1231231,
		}
		_, _ = uuservice.UserRelationFriendListService(context.Background(), &req)
	}
}
func BenchmarkUserRelationFriendListServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := useruser.RelationFriendListRequest{
				Token:  "1231312",
				UserId: 1231231,
			}
			_, _ = uuservice.UserRelationFriendListService(context.Background(), &req)
		}
	})
}
