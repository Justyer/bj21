package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/pkg/enum"
	"fxkt.tech/bj21/internal/pkg/json"
	"google.golang.org/grpc"
)

const (
	host = "0.0.0.0:9009"
)

func initcli() (v1.BlackJack_LogicConnClient, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := v1.NewBlackJackClient(conn)
	logicCli, err := client.LogicConn(ctx)
	if err != nil {
		panic(err)
	}
	return logicCli, cancel
}

func TestLogic(t *testing.T) {
	logicCli, cancel := initcli()
	defer cancel()

	err := logicCli.Send(&v1.Message{
		Cmd: enum.CmdLogin,
		Text: json.ToBytes(&v1.LoginRequest{
			Name: "dengxc",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	var token string

	for {
		rmsg, err := logicCli.Recv()
		if err != nil {
			t.Fatal(err)
		}

		switch rmsg.Cmd {
		case enum.CmdLogin:
			var res v1.LoginReply
			json.ToObjectByBytes(rmsg.Text, &res)
			token = res.Token
			logicCli.Send(&v1.Message{
				Cmd:  "tablelist",
				Text: json.ToBytes(&v1.TableListRequest{Token: token}),
			})
		case enum.CmdTableList:
			var res v1.TableListReply
			json.ToObjectByBytes(rmsg.Text, &res)
			for _, tb := range res.Tables {
				fmt.Println(tb.String())
			}
			logicCli.Send(&v1.Message{
				Cmd: enum.CmdSitDown,
				Text: json.ToBytes(&v1.SitDownRequest{
					Token:    token,
					TableSeq: res.Tables[0].Seq,
				}),
			})
		case enum.CmdSitDown:
			var res v1.SitDownReply
			json.ToObjectByBytes(rmsg.Text, &res)
			fmt.Println("enter table:", res.Err)
			goto x
		}

		t.Log(rmsg.Cmd, string(rmsg.Text))
	}
x:
	logicCli.Send(&v1.Message{
		Cmd:  enum.CmdTableList,
		Text: json.ToBytes(&v1.TableListRequest{Token: token}),
	})
	rmsg, err := logicCli.Recv()
	if err != nil {
		t.Fatal(err)
	}
	var res v1.TableListReply
	json.ToObjectByBytes(rmsg.Text, &res)
	for _, tb := range res.Tables {
		fmt.Println(tb.String())
	}
}

func TestLogin(t *testing.T) {
	logicCli, cancel := initcli()
	defer cancel()
	err := logicCli.Send(&v1.Message{
		Cmd: enum.CmdLogin,
		Text: json.ToBytes(&v1.LoginRequest{
			Name: "dengxc",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	rmsg, err := logicCli.Recv()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rmsg.Cmd, string(rmsg.Text))
}

func TestTableList(t *testing.T) {
	logicCli, cancel := initcli()
	defer cancel()
	err := logicCli.Send(&v1.Message{
		Cmd: "tablelist",
	})
	if err != nil {
		t.Fatal(err)
	}

	rmsg, err := logicCli.Recv()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rmsg.Cmd, string(rmsg.Text))
}
