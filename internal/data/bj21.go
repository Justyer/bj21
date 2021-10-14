package data

import (
	"context"
	"strconv"

	"fxkt.tech/bj21/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type bj21Repo struct {
	data *Data
	log  *log.Helper
}

func Newbj21Repo(data *Data, logger log.Logger) biz.BlackJackRepo {
	return &bj21Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *bj21Repo) EnterRoom(ctx context.Context, roomid int32) (token string, err error) {
	return strconv.FormatInt(int64(roomid), 10), nil
}
