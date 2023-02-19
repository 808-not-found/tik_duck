package userplatservice_test

import (
	"context"
	"reflect"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

func TestUserCommentActionService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *userplat.CommentActionRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *userplat.CommentActionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userplatservice.UserCommentActionService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCommentActionService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCommentActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
