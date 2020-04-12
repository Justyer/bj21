package bj21

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Stage struct {
	ps     []*Poker
	p1, p2 *Player
}

func (s *Stage) InitStage() {
	s.InitPokers()
	s.Deal(1)
	s.Deal(2)
	s.Deal(1)
	s.Deal(2)
}

// 设置玩家1
func (s *Stage) SetP1(p *Player) {
	s.p1 = p
}

// 设置玩家2
func (s *Stage) SetP2(p *Player) {
	s.p2 = p
}

// 初始化牌库，俗称洗牌
func (s *Stage) InitPokers() {
	for i := 1; i <= 11; i++ {
		s.ps = append(s.ps, &Poker{Number: i, Uni: i})
	}
	for i := 0; i < 11; i++ {
		x := rand.Intn(11)
		s.ps[i], s.ps[x] = s.ps[x], s.ps[i]
	}
}

// 查看牌库情况
func (s *Stage) ShowRepo() {
	var sls []string
	for _, p := range s.ps {
		sls = append(sls, p.String())
	}
	fmt.Println("repo->", strings.Join(sls, "\t"))

	if s.p1 != nil {
		fmt.Println(s.p1.String())
	}
	if s.p2 != nil {
		fmt.Println(s.p2.String())
	}
}

// 发牌
func (s *Stage) Deal(pnum int32) {
	p := s.ps[len(s.ps)-1]
	s.ps = append(s.ps[:len(s.ps)-1])
	switch pnum {
	case 1:
		s.p1.ps = append(s.p1.ps, p)
	case 2:
		s.p2.ps = append(s.p2.ps, p)
	}
}

func (s *Stage) Result(tnum int32) int32 {
	var p1n, p2n int
	for _, p := range s.p1.ps {
		p1n = p1n + p.Number
	}
	for _, p := range s.p2.ps {
		p2n = p2n + p.Number
	}
	switch {
	case p1n > 21 && p2n > 21:
		return 0
	case p1n <= 21 && p2n > 21:
		return 1
	case p1n > 21 && p2n <= 21:
		return 2
	case p1n == p2n:
		return 0
	case p1n > p2n:
		return 1
	case p1n < p2n:
		return 2
	}
	return -1
}
