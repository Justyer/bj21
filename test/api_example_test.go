package test

import (
	"bytes"
	"net/http"
	"net/http/httputil"
	"testing"

	v1 "fxkt.tech/bj21/api/bj21/v1"

	"fxkt.tech/bj21/internal/pkg/json"
)

const (
	host = "http://0.0.0.0:8008"
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
