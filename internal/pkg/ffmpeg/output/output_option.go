package output

import "fmt"

type OptionFunc func(*option)

type option struct {
	maps                  []string // mean is -map.
	cv, ca                string   // cv is c:v, ca is c:a.
	bv, ba                int64    // bv is b:v, ba is b:a.
	metadatas             []string // mean is -metadata.
	threads               int32    // thread counts, default 4.
	max_muxing_queue_size int32    // queue size when muxing, default 4086.
	movflags              string   // location of mp4 moov.
	ss                    int64    // ss is start_time.
	t                     int64    // t is duration.
	f                     string   // f is -f format.
	file                  string
}

func SetMap(m string) OptionFunc {
	return func(o *option) {
		o.maps = append(o.maps, m)
	}
}

func SetThread(n int32) OptionFunc {
	return func(o *option) {
		if n > 0 {
			o.threads = n
		}
	}
}

func SetVideoCoder(cv string) OptionFunc {
	return func(o *option) {
		o.cv = cv
	}
}

func SetAudioCoder(ca string) OptionFunc {
	return func(o *option) {
		o.ca = ca
	}
}

func SetVideoBitrate(bv int64) OptionFunc {
	return func(o *option) {
		o.bv = bv
	}
}

func SetAudioBitrate(ba int64) OptionFunc {
	return func(o *option) {
		o.ba = ba
	}
}

func SetMetadata(k, v string) OptionFunc {
	return func(o *option) {
		o.metadatas = append(o.metadatas, fmt.Sprintf("%s=%s", k, v))
	}
}

func SetStartTime(ss int64) OptionFunc {
	return func(o *option) {
		o.ss = ss
	}
}

func SetDuration(t int64) OptionFunc {
	return func(o *option) {
		o.t = t
	}
}

func SetFile(f string) OptionFunc {
	return func(o *option) {
		o.file = f
	}
}
