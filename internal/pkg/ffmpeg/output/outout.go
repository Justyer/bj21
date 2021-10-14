package output

import (
	"fmt"
	"strconv"
	"strings"
)

// Output is common output info.
type Output struct {
	opt *option
}

func New(opts ...OptionFunc) *Output {
	o := &option{
		threads:               4,
		max_muxing_queue_size: 4086,
		movflags:              "faststart",
		cv:                    "libx264",
		ca:                    "copy",
	}
	for _, opt := range opts {
		opt(o)
	}
	return &Output{
		opt: o,
	}
}

func (o *Output) Params() (params []string) {
	if len(o.opt.maps) != 0 {
		for _, m := range o.opt.maps {
			if strings.Contains(m, ":") {
				params = append(params, "-map", fmt.Sprintf("%s", m))
			} else {
				params = append(params, "-map", fmt.Sprintf("[%s]", m))
			}
		}
	}
	if o.opt.cv != "" {
		params = append(params, "-c:v", o.opt.cv)
	}
	if o.opt.ca != "" {
		params = append(params, "-c:a", o.opt.ca)
	}
	if o.opt.bv != 0 {
		params = append(params, "-b:v", strconv.FormatInt(o.opt.bv, 10))
	}
	if o.opt.ba != 0 {
		params = append(params, "-b:a", strconv.FormatInt(o.opt.ba, 10))
	}
	if len(o.opt.metadatas) != 0 {
		for _, m := range o.opt.metadatas {
			params = append(params, "-metadata", m)
		}
	}
	if o.opt.threads != 0 {
		params = append(params, "-threads", strconv.FormatInt(int64(o.opt.threads), 10))
	}
	if o.opt.max_muxing_queue_size != 0 {
		params = append(params, "-max_muxing_queue_size", strconv.FormatInt(int64(o.opt.max_muxing_queue_size), 10))
	}
	if o.opt.movflags != "" {
		params = append(params, "-movflags", o.opt.movflags)
	}
	if o.opt.ss != 0 {
		params = append(params, "-ss", strconv.FormatInt(o.opt.ss, 10))
	}
	if o.opt.t != 0 {
		params = append(params, "-t", strconv.FormatInt(o.opt.t, 10))
	}
	if o.opt.f != "" {
		params = append(params, "-f", o.opt.f)
	}
	if o.opt.file != "" {
		params = append(params, o.opt.file)
	}
	return
}

func (o *Output) String() string {
	return strings.Join(o.Params(), " ")
}
