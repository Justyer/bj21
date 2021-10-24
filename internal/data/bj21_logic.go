package data

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/data/logic"
	"fxkt.tech/bj21/internal/pkg/json"
)

var emptyjson = []byte("{}")

func (r *bj21Repo) login(txt []byte, srv v1.BlackJack_LogicConnServer) []byte {
	var req v1.LoginRequest
	json.ToObjectByBytes(txt, &req)
	player := logic.NewPlayer(req.Name, srv)
	r.data.world.AddOnlinePlayer(player)
	token := player.GetToken()
	return json.ToBytes(&v1.LoginReply{
		Name:  req.Name,
		Token: token,
	})
}

func (r *bj21Repo) logout(txt []byte) []byte {
	var req v1.LogoutRequest
	json.ToObjectByBytes(txt, &req)
	p := r.data.world.GetPlayerByToken(req.Token)
	var errstr string
	if p == nil {
		errstr = "player is not exist."
	} else {
		if p.InTable != nil {
			p.InTable.StandUp(req.Token)
		}
		r.data.world.DelOfflinePlayer(p)
	}
	return json.ToBytes(&v1.LogoutReply{Err: errstr})
}

func (r *bj21Repo) tablelist() []byte {
	tlist := r.data.world.TableList()
	rtables := make([]*v1.Table, len(tlist))
	for i, tb := range tlist {
		var p1, p2 *v1.Player
		if tb.P1 != nil {
			p1 = &v1.Player{Name: tb.P1.Name}
		}
		if tb.P2 != nil {
			p2 = &v1.Player{Name: tb.P2.Name}
		}
		rtables[i] = &v1.Table{
			Name: tb.Name,
			Seq:  tb.Seq,
			P1:   p1,
			P2:   p2,
		}
	}
	return json.ToBytes(&v1.TableListReply{Tables: rtables})
}

func (r *bj21Repo) sitdown(txt []byte) []byte {
	var req v1.SitDownRequest
	json.ToObjectByBytes(txt, &req)
	p := r.data.world.GetPlayerByToken(req.Token)
	tb := r.data.world.GetTableBySeq(req.TableSeq)
	var errstr string
	err := tb.SitDown(p)
	if err != nil {
		errstr = err.Error()
	} else {
		tb.Broadcast(&v1.Message{Cmd: "table-action", Text: emptyjson}, req.Token)
		r.log.Debugw("cmd", "table-action", "ocmd", "sitdown")
	}
	return json.ToBytes(&v1.SitDownReply{
		Table: tableLToP(tb, req.Token),
		Err:   errstr,
	})
}

func (r *bj21Repo) tableinfo(txt []byte) []byte {
	var req v1.TableInfoRequest
	json.ToObjectByBytes(txt, &req)
	tb := r.data.world.GetTableBySeq(req.TableSeq)
	return json.ToBytes(&v1.TableInfoReply{Table: tableLToP(tb, req.Token)})
}

func (r *bj21Repo) standup(txt []byte) []byte {
	var req v1.StandUpRequest
	json.ToObjectByBytes(txt, &req)
	tb := r.data.world.GetTableBySeq(req.TableSeq)
	var errstr string
	err := tb.StandUp(req.Token)
	if err != nil {
		errstr = err.Error()
	} else {
		tb.Broadcast(&v1.Message{Cmd: "table-action", Text: emptyjson}, req.Token)
		r.log.Debugw("cmd", "table-action", "ocmd", "standup")
	}
	return json.ToBytes(&v1.StandUpReply{Err: errstr})
}

func (r *bj21Repo) startgame(txt []byte) []byte {
	var req v1.StartGameRequest
	json.ToObjectByBytes(txt, &req)
	tb := r.data.world.GetTableBySeq(req.TableSeq)
	var errstr string
	if tb == nil {
		errstr = "table is not exist."
	} else {
		err := tb.StartGame()
		if err != nil {
			errstr = err.Error()
		} else {
			tb.Broadcast(&v1.Message{Cmd: "table-action", Text: emptyjson}, "")
			r.log.Debugw("cmd", "table-action", "ocmd", "startgame")
		}
	}
	return json.ToBytes(&v1.StartGameReply{Err: errstr})
}

func (r *bj21Repo) endGame(txt []byte) []byte {
	var req v1.EndGameRequest
	json.ToObjectByBytes(txt, &req)
	tb := r.data.world.GetTableBySeq(req.TableSeq)
	var errstr string
	if tb == nil {
		errstr = "table is not exist."
	} else {
		err := tb.EndGame()
		if err != nil {
			errstr = err.Error()
		} else {
			tb.Broadcast(&v1.Message{Cmd: "table-action", Text: emptyjson}, "")
			r.log.Debugw("cmd", "table-action", "ocmd", "endGame")
		}
	}
	return json.ToBytes(&v1.EndGameReply{Err: errstr})
}
