package input

type OptionFunc func(*option)

type option struct {
	i   string   // i is input file.
	ss  int64    // ss is start_time.
	t   int64    // t is duration.
	ext []string // extra params.
}

func SetI(i string) OptionFunc {
	return func(o *option) {
		o.i = i
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
