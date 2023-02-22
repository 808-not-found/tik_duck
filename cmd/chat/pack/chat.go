package pack

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/chat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/chat"
)

func Msgs(ctx context.Context, ms []*db.Message) (res []*chat.Message, err error) {
	return
}
