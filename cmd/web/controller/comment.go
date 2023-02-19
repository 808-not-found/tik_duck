package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/cloudwego/hertz/pkg/app"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.CommentActionRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserCommentAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.CommentListRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserCommentList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
