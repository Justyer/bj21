package logic

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"

	v1 "fxkt.tech/bj21/api/bj21/v1"
)

const (
	statusIdle   = "idle"
	statusGaming = "gaming"
)

type Table struct {
	Name   string
	Seq    string
	P1, P2 *Player
	Pk     Poker
	Status string
	round  int32 // 0: p1, 1: p2

	mu sync.Mutex
}

func NewTable(name string) *Table {
	return &Table{
		Name:   name,
		Status: statusIdle,
		Seq:    fmt.Sprintf("TB%s", defaulthid.GenId()),
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

// 本桌入座
func (t *Table) SitDown(p *Player) error {
	if t.Status != statusIdle {
		return errors.New("table is gaming.")
	}

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

// 起身离开本桌
func (t *Table) StandUp(token string) error {
	// 游戏中如果起身的话，直接判负
	if t.Status == statusGaming {
		err := t.EndGame()
		if err != nil {
			return err
		}
	}

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
	}
	if t.P2 != nil && t.P2.GetToken() != token {
		t.P2.TellMe(msg)
	}
}

func (t *Table) StartGame() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Status != statusIdle {
		return errors.New("table is playing.") // 游戏中不能开始游戏
	}
	if t.P1 == nil || t.P2 == nil {
		return errors.New("table player is not full.")
	}

	t.Pk = NewPoker()        // 一副新牌
	t.Pk.Wash()              // 洗牌
	t.round = rand.Int31n(2) // 确定先手
	t.P1.WashHands()         // 重置p1手牌
	t.P2.WashHands()         // 重置p2手牌
	fmt.Println("wash pokers result: ", t.Pk.String())
	for i := 0; i < 4; i++ {
		c := t.Pk.Deal()
		t.currentRoundPlayer().GiveMe(c) // 发牌
		t.switchRound()                  // 切换回合
	}
	t.Status = statusGaming
	return nil
}

func (t *Table) Hit(token string) error {
	cp := t.currentRoundPlayer()
	if cp.token != token {
		return errors.New("this is not your round.")
	}
	t.switchRound()
	return nil
}

func (t *Table) EndGame() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Status != statusGaming {
		return errors.New("table is not in playing.")
	}
	t.Status = statusIdle
	return nil
}

func (t *Table) currentRoundPlayer() *Player {
	if t.round == 0 {
		return t.P1
	} else if t.round == 1 {
		return t.P2
	}
	return nil
}

func (t *Table) switchRound() {
	t.round = 1 - t.round
}
