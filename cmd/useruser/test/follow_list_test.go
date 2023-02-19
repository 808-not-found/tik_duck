package uuservice_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

func TestUserRelationFollowListService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *useruser.RelationFollowListRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *useruser.RelationFollowListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uuservice.UserRelationFollowListService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRelationFollowListService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRelationFollowListService() = %v, want %v", got, tt.want)
			}
		})
	}
}
