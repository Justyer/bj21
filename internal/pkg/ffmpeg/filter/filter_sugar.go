package filter

import (
	"fmt"
	"strings"

	"fxkt.tech/bj21/internal/pkg/ffmpeg/math"
)

type Stream string

const (
	StreamAudio Stream = "a"
	StreamVideo Stream = "v"
)

func SelectStream(idx int, s Stream, must bool) string {
	var qm string
	if !must {
		qm = "?"
	}
	return fmt.Sprintf("%d:%s%s", idx, s, qm)
}

type LogoPos string

const (
	LogoTopLeft     LogoPos = "TopLeft"
	LogoTopRight    LogoPos = "TopRight"
	LogoBottomRight LogoPos = "BottomRight"
	LogoBottomLeft  LogoPos = "BottomLeft"
)

func Logo(dx, dy int64, pos LogoPos) string {
	switch pos {
	case LogoTopLeft:
		return fmt.Sprintf("overlay=%d:y=%d", dx, dy)
	case LogoTopRight:
		return fmt.Sprintf("overlay=W-w-%d:y=%d", dx, dy)
	case LogoBottomRight:
		return fmt.Sprintf("overlay=W-w-%d:y=H-h-%d", dx, dy)
	case LogoBottomLeft:
		return fmt.Sprintf("overlay=%d:y=H-h-%d", dx, dy)
	}
	return ""
}

// func Overlay(x, y int64, eofaction string) string {

// }

func Scale(w, h int64) string {
	return fmt.Sprintf("scale=%d:%d",
		math.CeilOdd(w),
		math.CeilOdd(h),
	)
}

func Split(n int) string {
	return fmt.Sprintf("split=%d", n)
}

func Trim(s, e int64) string {
	var ps []string
	if s != 0 {
		ps = append(ps, fmt.Sprintf("start=%d", s))
	}
	if e != 0 {
		ps = append(ps, fmt.Sprintf("end=%d", e))
	}
	var eqs string
	var psstr string = strings.Join(ps, ":")
	if psstr != "" {
		eqs = "="
	}
	return fmt.Sprintf("trim%s%s", eqs, psstr)
}

func Delogo(x, y, w, h int64) string {
	return fmt.Sprintf("delogo=%d:%d:%d:%d",
		x+1, y+1, w-2, h-2,
	)
}
