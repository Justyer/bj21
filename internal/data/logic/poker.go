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
	pk []*Card
}

type Card struct {
	num int
}

type Poker []*Card

func NewPoker() Poker {
	ps := make(Poker, 10)
	for i := 1; i <= 10; i++ {
		ps[i-1] = &Card{
			num: i,
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
	var nums []int
	for _, pk := range p {
		nums = append(nums, pk.num)
	}
	return fmt.Sprintf("%v", nums)
}
