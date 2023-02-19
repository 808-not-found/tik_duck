package userservice_test

import (
	"context"
	"reflect"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func TestUserPublishListService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *user.PublishListRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.PublishListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userservice.UserPublishListService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPublishListService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPublishListService() = %v, want %v", got, tt.want)
			}
		})
	}
}
