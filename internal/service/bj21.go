package service

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type BlackJackService struct {
	v1.UnimplementedBlackJackServer

	uc  *biz.BlackJackUsecase
	log *log.Helper
}

func NewBlackJackService(uc *biz.BlackJackUsecase, logger log.Logger) *BlackJackService {
	return &BlackJackService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BlackJackService) LogicConn(srv v1.BlackJack_LogicConnServer) (err error) {
	return s.uc.LogicConn(srv)
}
