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

		var rtxt []byte
		switch msg.Cmd {
		case enum.CmdLogin:
			rtxt = r.login(msg.Text, srv)
		case enum.CmdLogout:
			rtxt = r.logout(msg.Text)
		case enum.CmdTableList:
			rtxt = r.tablelist()
		case enum.CmdSitDown:
			rtxt = r.sitdown(msg.Text)
		case enum.CmdTableInfo:
			rtxt = r.tableinfo(msg.Text)
		case enum.CmdStandUp:
			rtxt = r.standup(msg.Text)
		}
		srv.Send(&v1.Message{Cmd: msg.Cmd, Text: rtxt})
		r.log.Debugw("cmd", msg.Cmd, "req", string(msg.Text), "res", string(rtxt))
	}
}
