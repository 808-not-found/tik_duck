package userservice

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func TestUserRegisterService(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.UserRegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		want1   string
		want2   int64
		want3   string
		wantErr bool
	}{
		{"1" , args{}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := UserRegisterService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegisterService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserRegisterService() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UserRegisterService() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("UserRegisterService() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("UserRegisterService() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
