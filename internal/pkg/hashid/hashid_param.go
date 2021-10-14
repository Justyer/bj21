package hashid

import (
	"time"
)

type HashIdParam struct {
	NowTime int64
	Rand    int64
}

func (p *HashIdParam) GetYearMonth() int64 {
	tm := time.Unix(0, p.NowTime)
	return int64(tm.Year()*100 + int(tm.Month()))
}
