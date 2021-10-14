package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type BlackJackRepo interface {
	EnterRoom(ctx context.Context, roomid int32) (token string, err error)
}

type BlackJackUsecase struct {
	repo BlackJackRepo
	log  *log.Helper
}

func NewBlackJackUsecase(repo BlackJackRepo, logger log.Logger) *BlackJackUsecase {
	return &BlackJackUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BlackJackUsecase) EnterRoom(ctx context.Context, roomid int32) (token string, err error) {
	return uc.repo.EnterRoom(ctx, roomid)
}
