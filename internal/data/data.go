package data

import (
	"fxkt.tech/bj21/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var (
	ProviderSet = wire.NewSet(NewData, Newbj21Repo)
)

type Data struct {
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelWarn, "status", "closing the data resources")
	}

	return &Data{}, cleanup, nil
}
