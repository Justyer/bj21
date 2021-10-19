package logic

import (
	"fmt"

	v1 "fxkt.tech/bj21/api/bj21/v1"
	"github.com/google/uuid"
)

type Player struct {
	Name  string
	token string
	conn  v1.BlackJack_LogicConnServer
}

func NewPlayer(name string, conn v1.BlackJack_LogicConnServer) *Player {
	return &Player{
		Name:  name,
		token: uuid.NewString(),
		conn:  conn,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Name: %s,Token: %s", p.Name, p.token)
}

func (p *Player) GetToken() string {
	return p.token
}

func (p *Player) TellMe(msg *v1.Message) error {
	return p.conn.Send(msg)
}
