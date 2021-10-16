package logic

import (
	"errors"
	"fmt"
	"sync"
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
	return nil
}

// 起身离开本桌
func (t *Table) StandUp(p *Player) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.P1.Name == p.Name {
		t.P1 = nil
	} else if t.P2.Name == p.Name {
		t.P2 = nil
	} else {
		return errors.New("you are not in this table.")
	}
	return nil
}
