// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	douyin_user "github.com/808-not-found/tik_duck/cmd/user/kitex_gen/douyin_user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	User_Register(ctx context.Context, Req *douyin_user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserRegisterResponse, err error)
	User_GetFeed(ctx context.Context, Req *douyin_user.DouyinFeedRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinFeedResponse, err error)
	User_Login(ctx context.Context, Req *douyin_user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserLoginResponse, err error)
	User_Info(ctx context.Context, Req *douyin_user.DouyinUserInfoRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserInfoResponse, err error)
	User_PublishList(ctx context.Context, Req *douyin_user.DouyinPublishListRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinPublishListResponse, err error)
	User_PublishAction(ctx context.Context, Req *douyin_user.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinPublishActionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) User_Register(ctx context.Context, Req *douyin_user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_Register(ctx, Req)
}

func (p *kUserServiceClient) User_GetFeed(ctx context.Context, Req *douyin_user.DouyinFeedRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_GetFeed(ctx, Req)
}

func (p *kUserServiceClient) User_Login(ctx context.Context, Req *douyin_user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_Login(ctx, Req)
}

func (p *kUserServiceClient) User_Info(ctx context.Context, Req *douyin_user.DouyinUserInfoRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_Info(ctx, Req)
}

func (p *kUserServiceClient) User_PublishList(ctx context.Context, Req *douyin_user.DouyinPublishListRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_PublishList(ctx, Req)
}

func (p *kUserServiceClient) User_PublishAction(ctx context.Context, Req *douyin_user.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *douyin_user.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.User_PublishAction(ctx, Req)
}