package data

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/data/logic"
	"fxkt.tech/bj21/internal/pkg/json"
)

func (r *bj21Repo) login(txt []byte, srv v1.BlackJack_LogicConnServer) []byte {
	var req v1.LoginRequest
	json.ToObjectByBytes(txt, &req)
	player := logic.NewPlayer(req.Name, srv)
	r.data.world.AddOnlinePlayers(player)
	token := player.GetToken()
	res := &v1.LoginReply{Token: token}
	return json.ToBytes(res)
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
	res := &v1.TableListReply{Tables: rtables}
	return json.ToBytes(res)
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
	}
	res := &v1.SitDownReply{Err: errstr}
	return json.ToBytes(res)
}
