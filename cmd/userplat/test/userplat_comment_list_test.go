package userplatservice_test

import (
	"context"
	"reflect"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

func TestUserCommentListService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *userplat.CommentListRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *userplat.CommentListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userplatservice.UserCommentListService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCommentListService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCommentListService() = %v, want %v", got, tt.want)
			}
		})
	}
}
