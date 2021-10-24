package data

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/data/logic"
)

func tableLToP(l *logic.Table, token string) *v1.Table {
	if l == nil {
		return nil
	}
	return &v1.Table{
		Name:   l.Name,
		Seq:    l.Seq,
		P1:     playerLToP(l.P1),
		P2:     playerLToP(l.P2),
		Me:     l.MyLoc(token),
		Status: l.Status,
	}
}

func playerLToP(l *logic.Player) *v1.Player {
	if l == nil {
		return nil
	}
	var cs []*v1.Card
	for _, c := range l.Hands.Cards() {
		cs = append(cs, cardLToP(c))
	}
	return &v1.Player{
		Name:  l.Name,
		Cards: cs,
	}
}

func cardLToP(c *logic.Card) *v1.Card {
	if c == nil {
		return nil
	}
	return &v1.Card{
		Num: c.Num,
		Id:  c.Id,
	}
}
