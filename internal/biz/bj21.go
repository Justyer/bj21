package biz

import (
	v1 "fxkt.tech/bj21/api/bj21/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type BlackJackRepo interface {
	LogicConn(srv v1.BlackJack_LogicConnServer) error
}

type BlackJackUsecase struct {
	repo BlackJackRepo
	log  *log.Helper
}

func NewBlackJackUsecase(repo BlackJackRepo, logger log.Logger) *BlackJackUsecase {
	return &BlackJackUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BlackJackUsecase) LogicConn(srv v1.BlackJack_LogicConnServer) error {
	return uc.repo.LogicConn(srv)
}
