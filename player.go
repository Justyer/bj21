package bj21

import (
	"fmt"
	"strings"
)

type Player struct {
	Number int
	ps     []*Poker
}

func (p *Player) String() string {
	var sls []string
	var total int
	for _, p := range p.ps {
		sls = append(sls, p.String())
		total = total + p.Number
	}
	return fmt.Sprintf("p%d-> %s | %2d/21", p.Number, strings.Join(sls, "\t"), total)
}
