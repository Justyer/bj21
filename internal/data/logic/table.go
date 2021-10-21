package logic

import (
	"errors"
	"fmt"
	"sync"

	v1 "fxkt.tech/bj21/api/bj21/v1"
)

type Table struct {
	Name   string
	Seq    string
	P1, P2 *Player
	Pk     *Poker

	mu sync.Mutex
}

func NewTable(name string) *Table {
	return &Table{
		Name: name,
		Seq:  fmt.Sprintf("TB%s", defaulthid.GenId()),
	}
}

func (t *Table) String() string {
	var p1str, p2str string
	if t.P1 != nil {
		p1str = t.P1.String()
	}
	if t.P2 != nil {
		p2str = t.P2.String()
	}
	return fmt.Sprintf("Name:\t%s\nPlayer:\t%s\t%s\n", t.Seq, p1str, p2str)
}

// 本桌入座
func (t *Table) SitDown(p *Player) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.P1 == nil {
		t.P1 = p
	} else if t.P2 == nil {
		t.P2 = p
	} else {
		return errors.New("sits of table is full.")
	}
	p.InTable = t
	return nil
}

// 我的位置： p1 or p2.
func (t *Table) MyLoc(token string) int32 {
	if t.P1 != nil && t.P1.token == token {
		return 1
	} else if t.P2 != nil && t.P2.token == token {
		return 2
	} else {
		return 0
	}
}

// 起身离开本桌
func (t *Table) StandUp(token string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.P1 != nil && t.P1.token == token {
		t.P1 = nil
	} else if t.P2 != nil && t.P2.token == token {
		t.P2 = nil
	} else {
		return errors.New("you are not in this table.")
	}
	return nil
}

// 房间内广播
func (t *Table) Broadcast(msg *v1.Message, token string) {
	if t.P1 != nil && t.P1.GetToken() != token {
		t.P1.TellMe(msg)
	} else if t.P2 != nil && t.P2.GetToken() != token {
		t.P2.TellMe(msg)
	}
}
