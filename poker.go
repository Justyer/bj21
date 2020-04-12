package bj21

import "fmt"

type Poker struct {
	Number int
	Uni    int
}

func (p *Poker) String() string {
	return fmt.Sprintf("%2d:#%x", p.Number, p.Uni)
}
