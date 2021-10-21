package logic

import (
	"errors"

	"fxkt.tech/bj21/internal/pkg/hashid"
)

var defaulthid = hashid.New("bj21", 16)

type World struct {
	tables  []*Table
	players map[string]*Player // number of player online, kv=token:player
}

func NewWorld() *World {
	w := &World{
		players: make(map[string]*Player),
	}
	w.init()
	return w
}

func (w *World) init() {
	// 系统自带房间
	w.tables = append(w.tables,
		NewTable("BEDROOM"),
		NewTable("HOSPITAL"),
	)
}

func (w *World) TableList() []*Table {
	return w.tables
}

func (w *World) GetTable(seq string) *Table {
	for i, table := range w.tables {
		if table.Seq == seq {
			return w.tables[i]
		}
	}
	return nil
}

// 将上线的玩家放入在线玩家池中
func (w *World) AddOnlinePlayer(p *Player) error {
	if p == nil {
		return errors.New("player can not empty.")
	}
	w.players[p.GetToken()] = p
	return nil
}

func (w *World) DelOfflinePlayer(p *Player) error {
	if p == nil {
		return errors.New("player can not empty.")
	}
	token := p.GetToken()
	if _, ok := w.players[token]; ok {
		delete(w.players, token)
		return nil
	}
	return errors.New("player is not exist.")
}

func (w *World) GetTableBySeq(seq string) *Table {
	for i, tb := range w.tables {
		if tb.Seq == seq {
			return w.tables[i]
		}
	}
	return nil
}

func (w *World) GetPlayerByToken(token string) *Player {
	if p, ok := w.players[token]; ok {
		return p
	}
	return nil
}
