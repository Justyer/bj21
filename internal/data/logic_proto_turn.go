package data

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/data/logic"
)

func tableLToP(l *logic.Table) *v1.Table {
	if l == nil {
		return nil
	}
	return &v1.Table{
		Name: l.Name,
		Seq:  l.Seq,
		P1:   playerLToP(l.P1),
		P2:   playerLToP(l.P2),
	}
}

func playerLToP(l *logic.Player) *v1.Player {
	if l == nil {
		return nil
	}
	return &v1.Player{
		Name: l.Name,
	}
}
