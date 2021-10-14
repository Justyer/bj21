package ffmpeg

import (
	"context"
	"testing"

	"fxkt.tech/bj21/internal/pkg/ffmpeg/filter"
	"fxkt.tech/bj21/internal/pkg/ffmpeg/input"
	"fxkt.tech/bj21/internal/pkg/ffmpeg/output"
)

func TestFFmpeg(t *testing.T) {
	ff := Default()
	ff.AddInput(
		input.New(
			input.SetI("vieo.mp4"),
		),
	)
	ff.AddFilter(
		filter.New(
			filter.SetInStream("0"),
			filter.SetContent("scale=trunc(oh*a/2)*2:720"),
			filter.SetOutStream("x720"),
		),
		filter.New(
			filter.SetInStream("x720"),
			filter.SetContent("delogo=0:0:100:100"),
			filter.SetOutStream("xx720"),
		),
	)
	ff.AddOutput(
		output.New(
			output.SetMap("xx720"),
			output.SetMap("0:a"),
			output.SetMetadata("comment", "yilan888"),
			output.SetFile("out720.mp4"),
		),
	)
	ff.Run(context.Background())
	t.Log(ff.Sentence)
}
