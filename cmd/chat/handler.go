package main

import (
	"context"

	chat "github.com/808-not-found/tik_duck/kitex_gen/chat"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// GetMsg implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetMsg(
	ctx context.Context,
	req *chat.MessageChatRequest,
) (resp *chat.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// PostMsg implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) PostMsg(
	ctx context.Context,
	req *chat.RelationActionRequest,
) (resp *chat.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
