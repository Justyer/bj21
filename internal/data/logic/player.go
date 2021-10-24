package logic

import (
	"fmt"

	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/pkg/json"
	"github.com/google/uuid"
)

type Player struct {
	Name    string
	token   string
	conn    v1.BlackJack_LogicConnServer
	InTable *Table
	Hands   *Hands
}

func NewPlayer(name string, conn v1.BlackJack_LogicConnServer) *Player {
	return &Player{
		Name:  name,
		token: uuid.NewString(),
		conn:  conn,
		Hands: NewHands(),
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Name: %s,Token: %s", p.Name, p.token)
}

func (p *Player) GetToken() string {
	return p.token
}

func (p *Player) TellMe(msg *v1.Message) error {
	fmt.Println("TellMe", p.Name, json.ToString(msg))
	return p.conn.Send(msg)
}

func (p *Player) GiveMe(c *Card) {
	p.Hands.Insert(c)
}

func (p *Player) WashHands() {
	p.Hands = NewHands()
}
