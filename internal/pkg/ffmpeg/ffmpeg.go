package ffmpeg

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"strings"

	"fxkt.tech/bj21/internal/pkg/ffmpeg/filter"
	"fxkt.tech/bj21/internal/pkg/ffmpeg/input"
	"fxkt.tech/bj21/internal/pkg/ffmpeg/output"
)

type FFmpeg struct {
	cmd      string
	v        string // v is log level.
	y        bool   // y is yes for overwrite output file.
	inputs   input.Inputs
	filters  filter.Filters
	outputs  output.Outputs
	Sentence string
}

func Default() *FFmpeg {
	return &FFmpeg{
		cmd: "ffmpeg",
		y:   true,
	}
}

func (ff *FFmpeg) CmdLoc(loc string) {
	ff.cmd = loc
}

func (ff *FFmpeg) Yes(y bool) {
	ff.y = y
}

func (ff *FFmpeg) LogLevel(v string) {
	ff.v = v
}

func (ff *FFmpeg) AddInput(inputs ...*input.Input) {
	ff.inputs = append(ff.inputs, inputs...)
}

func (ff *FFmpeg) AddFilter(filters ...*filter.Filter) {
	ff.filters = append(ff.filters, filters...)
}

func (ff *FFmpeg) AddOutput(outputs ...*output.Output) {
	ff.outputs = append(ff.outputs, outputs...)
}

func (ff *FFmpeg) Params() (params []string) {
	if ff.v != "" {
		params = append(params, "-v", ff.v)
	}
	if ff.y {
		params = append(params, "-y")
	}
	params = append(params, ff.inputs.Params()...)
	params = append(params, ff.filters.Params()...)
	params = append(params, ff.outputs.Params()...)
	return
}

func (ff *FFmpeg) Run(ctx context.Context) (err error) {
	cc := exec.CommandContext(ctx, ff.cmd, ff.Params()...)
	ff.Sentence = cc.String()
	retbytes, err := cc.CombinedOutput()
	log.Println(string(retbytes))
	if err != nil {
		err = errors.New(strings.ReplaceAll(string(retbytes), "\n", "|"))
	}
	return
}
