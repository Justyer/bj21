package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"

	v1 "fxkt.tech/bj21/api/bj21/v1"
	"google.golang.org/grpc"

	"fxkt.tech/bj21/internal/pkg/json"
)

const (
	host     = "http://0.0.0.0:8008"
	grpchost = "0.0.0.0:9009"
)

func TestEnterRoom(t *testing.T) {
	req := v1.EnterRoomRequest{RoomId: 8989}
	reqb := json.ToBytes(&req)
	url := host + "/bj21.v1.BlackJack/EnterRoom"
	res, err := http.Post(url, "application/json", bytes.NewReader(reqb))
	if err != nil {
		t.Error(err)
	}
	_ = res
	qb, _ := httputil.DumpRequest(res.Request, true)
	t.Log(string(qb))
	sb, _ := httputil.DumpResponse(res, true)
	t.Log(string(sb))
}

func TestChat(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	conn, err := grpc.DialContext(ctx, grpchost, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	client := v1.NewBlackJackClient(conn)
	cli, err := client.Chat(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 9; i++ {
		err = cli.Send(&v1.ChatContent{Message: fmt.Sprintf("hello: %d", i)})
		if err != nil {
			t.Fatal(err)
		}
	}
}
