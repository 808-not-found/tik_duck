package userservice

import (
	"context"
	"reflect"
	"testing"

	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func TestUserInfoService(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.UserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		want1   string
		want2   *user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := UserInfoService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInfoService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserInfoService() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UserInfoService() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("UserInfoService() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
