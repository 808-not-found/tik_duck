package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/808-not-found/tik_duck/cmd/web/controller"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF { //nolint:all
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = controller.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserID, event.ToUserID)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserID, event.UserID)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserID)
			continue
		}

		pushEvent := controller.MessagePushEvent{
			FromUserID: event.UserID,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent) //nolint:all
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}
