package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

var tempChat = map[string][]Message{} //nolint:all

var messageIdSequence = int64(1) //nolint:all

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid.
func MessageAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserID := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := usersLoginInfo[token]; exist {
		userIDB, _ := strconv.Atoi(toUserID)
		chatKey := genChatKey(user.ID, int64(userIDB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := Message{
			ID:         messageIdSequence,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []Message{curMessage}
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list.
func MessageChat(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserID := c.Query("to_user_id")

	if user, exist := usersLoginInfo[token]; exist {
		userIDB, _ := strconv.Atoi(toUserID)
		chatKey := genChatKey(user.ID, int64(userIDB))

		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIDA int64, userIDB int64) string {
	if userIDA > userIDB {
		return fmt.Sprintf("%d_%d", userIDB, userIDA)
	}
	return fmt.Sprintf("%d_%d", userIDA, userIDB)
}
