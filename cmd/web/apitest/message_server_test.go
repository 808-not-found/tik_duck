package api_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/web/controller"
)

func TestMessageServer(t *testing.T) {
	e := newExpect(t)
	userIDA, _ := getTestUserToken(testUserA, e)
	userIDB, _ := getTestUserToken(testUserB, e)

	connA, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Connect server failed: %v\n", err)
		return
	}
	connB, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Connect server failed: %v\n", err)
		return
	}

	createChat(userIDA, connA, userIDB, connB)

	go readMessage(connB)
	sendMessage(userIDA, userIDB, connA)
}

func readMessage(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = controller.MessagePushEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Read message：%+v\n", event)
	}
}

func sendMessage(fromUserID int, toUserID int, fromConn net.Conn) {
	defer fromConn.Close()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		sendEvent := controller.MessageSendEvent{
			UserID:     int64(fromUserID),
			ToUserID:   int64(toUserID),
			MsgContent: "Test Content",
		}
		data, _ := json.Marshal(sendEvent)
		_, err := fromConn.Write(data)
		if err != nil {
			fmt.Printf("Send message failed: %v\n", err)
			return
		}
	}
	time.Sleep(time.Second)
}

func createChat(userIDA int, connA net.Conn, userIDB int, connB net.Conn) {
	chatEventA := controller.MessageSendEvent{
		UserID:   int64(userIDA),
		ToUserID: int64(userIDB),
	}
	chatEventB := controller.MessageSendEvent{
		UserID:   int64(userIDB),
		ToUserID: int64(userIDA),
	}
	eventA, _ := json.Marshal(chatEventA)
	eventB, _ := json.Marshal(chatEventB)
	_, err := connA.Write(eventA)
	if err != nil {
		fmt.Printf("Create chatA failed: %v\n", err)
		return
	}
	_, err = connB.Write(eventB)
	if err != nil {
		fmt.Printf("Create chatB failed: %v\n", err)
		return
	}
}
