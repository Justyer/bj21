package service

import (
	"context"

	v1 "fxkt.tech/bj21/api/bj21/v1"
	"fxkt.tech/bj21/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// BlackJackService implements service->LocalMediaService in lms.proto
type BlackJackService struct {
	v1.UnimplementedBlackJackServer

	uc  *biz.BlackJackUsecase
	log *log.Helper
}

func NewBlackJackService(uc *biz.BlackJackUsecase, logger log.Logger) *BlackJackService {
	return &BlackJackService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BlackJackService) EnterRoom(ctx context.Context, req *v1.EnterRoomRequest) (*v1.EnterRoomReply, error) {
	token, err := s.uc.EnterRoom(ctx, req.RoomId)
	if err != nil {
		return nil, err
	}

	return &v1.EnterRoomReply{RoomToken: token}, nil
}
