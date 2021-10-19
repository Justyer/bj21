package data

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/biz"
	"fxkt.tech/bj21/internal/pkg/enum"
	"github.com/go-kratos/kratos/v2/log"
)

type bj21Repo struct {
	data *Data
	log  *log.Helper
}

func Newbj21Repo(data *Data, logger log.Logger) biz.BlackJackRepo {
	return &bj21Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *bj21Repo) LogicConn(srv v1.BlackJack_LogicConnServer) error {
	for {
		msg, err := srv.Recv()
		if err != nil {
			return err
		}

		r.log.Debugw("rr", "req", "cmd", msg.Cmd, "text", string(msg.Text))

		var rmsg *v1.Message
		switch msg.Cmd {
		case enum.CmdLogin:
			rmsg = &v1.Message{
				Cmd:  msg.Cmd,
				Text: r.login(msg.Text, srv),
			}
		case enum.CmdTableList:
			rmsg = &v1.Message{
				Cmd:  msg.Cmd,
				Text: r.tablelist(),
			}
		case enum.CmdSitDown:
			rmsg = &v1.Message{
				Cmd:  msg.Cmd,
				Text: r.sitdown(msg.Text),
			}
		case enum.CmdTableInfo:
			rmsg = &v1.Message{
				Cmd:  msg.Cmd,
				Text: r.tableinfo(msg.Text),
			}
		}
		srv.Send(rmsg)
		r.log.Debugw("rr", "res", "cmd", rmsg.Cmd, "text", string(rmsg.Text))
	}
}
