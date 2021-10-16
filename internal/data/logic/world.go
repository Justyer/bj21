package logic

import "fxkt.tech/bj21/internal/pkg/hashid"

var defaulthid = hashid.New("bj21", 16)

type World struct {
	tables []*Table
	ps     []*Player // number of player online.
}

func NewWorld() *World {
	w := &World{}
	w.init()
	return w
}

func (w *World) init() {
	// 系统自带房间
	w.tables = append(w.tables, NewTable("system1"), NewTable("system2"))
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
func (w *World) AddOnlinePlayers(p *Player) error {
	w.ps = append(w.ps, p)
	return nil
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
	for i, p := range w.ps {
		if p.token == token {
			return w.ps[i]
		}
	}
	return nil
}
