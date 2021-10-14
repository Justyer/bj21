package hashid

import (
	"errors"
	"math/rand"
	"time"

	"github.com/speps/go-hashids"
)

type HashId struct {
	hid *hashids.HashID
}

func New(salt string, minlen int) *HashId {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minlen
	hid := hashids.NewWithData(hd)
	return &HashId{hid: hid}
}

// 生成jobid
// 每个id存储了两个值：纳秒时间戳、0～10000的随机数
// 理论上请求速度不超过 10000次/ns 就不会出现重复
// （实际上达到这个速度，系统就已经先挂了。。。）
func (h *HashId) GenId() string {
	tsn := time.Now().UnixNano()
	rdm := rand.Int63n(10000)
	hid, _ := h.hid.EncodeInt64([]int64{tsn, rdm})
	return hid
}

func (h *HashId) ParseId(hid string) (*HashIdParam, error) {
	params := h.hid.DecodeInt64(hid)
	if len(params) < 2 {
		return nil, errors.New("hid is not belong this system.")
	}
	return &HashIdParam{
		NowTime: params[0],
		Rand:    params[1],
	}, nil
}
