package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

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
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	feedReq.ActionType = int32(actionType)
	cid, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	ct := c.Query("comment_text")
	feedReq.CommentId = &cid
	feedReq.CommentText = &ct
	feedReq.Token = c.Query("token")
	feedReq.VideoId, _ = strconv.ParseInt(c.Query("video_id"), 10, 64)

	resp, err := rpc.UserCommentAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.CommentListRequest
	feedReq.Token = c.Query("token")
	feedReq.VideoId, _ = strconv.ParseInt(c.Query("video_id"), 10, 64)
	resp, err := rpc.UserCommentList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
