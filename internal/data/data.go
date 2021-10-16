package data

import (
	"fxkt.tech/bj21/internal/conf"
	"fxkt.tech/bj21/internal/data/logic"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var (
	ProviderSet = wire.NewSet(NewData, Newbj21Repo)
)

type Data struct {
	world *logic.World
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelWarn, "status", "closing the data resources")
	}

	world := logic.NewWorld()

	return &Data{
		world: world,
	}, cleanup, nil
}
