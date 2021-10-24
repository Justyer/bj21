package logic

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Hands struct {
	cs []*Card
}

func NewHands() *Hands {
	return &Hands{
		cs: make([]*Card, 0),
	}
}

func (h *Hands) Insert(c *Card) {
	h.cs = append(h.cs, c)
}

func (h *Hands) Cards() []*Card {
	return h.cs
}

type Card struct {
	Num int32
	Id  string
}

type Poker []*Card

func NewPoker() Poker {
	ps := make(Poker, 11)
	for i := 1; i <= 11; i++ {
		ps[i-1] = &Card{
			Num: int32(i),
			Id:  fmt.Sprintf("CD%s", defaulthid.GenId()),
		}
	}
	return ps
}

func (p *Poker) Wash() {
	for i := 0; i < 10; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		(*p)[x], (*p)[y] = (*p)[y], (*p)[x]
	}
}

func (p Poker) String() string {
	var nums []int32
	for _, pk := range p {
		nums = append(nums, pk.Num)
	}
	return fmt.Sprintf("%v", nums)
}

func (p *Poker) Deal() *Card {
	c := (*p)[len(*p)-1]
	(*p) = (*p)[:len((*p))-1]
	return c
}
