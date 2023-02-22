package pack

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/chat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/chat"
)

func Msg(ctx context.Context, m *db.Message) (res *chat.Message, err error) {
	t := m.CreatedAt.String()
	return &chat.Message{
		Id:         m.ID,
		ToUserId:   m.ToUserID,
		FromUserId: m.FromUserID,
		Content:    m.Content,
		CreateTime: &t,
	}, nil
}

func Msgs(ctx context.Context, ms []*db.Message) (res []*chat.Message, err error) {
	msgs := make([]*chat.Message, 0)
	for _, m := range ms {
		n, err := Msg(ctx, m)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, n)
	}
	return msgs, nil
}
