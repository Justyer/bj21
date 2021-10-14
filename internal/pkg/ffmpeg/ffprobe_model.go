package ffmpeg

type Probe struct {
	Streams []*ProbeStream `json:"streams,omitempty"`
	Format  *ProbeFormat   `json:"format,omitempty"`
}

type ProbeStream struct {
	Index      int64  `json:"index,omitempty"`
	CodecType  string `json:"codec_type,omitempty"`
	CodecName  string `json:"codec_name,omitempty"`
	Profile    string `json:"profile,omitempty"`
	Width      int64  `json:"width,omitempty"`
	Height     int64  `json:"height,omitempty"`
	PixFmt     string `json:"pix_fmt,omitempty"`
	Level      int64  `json:"level,omitempty"`
	RFrameRate string `json:"r_frame_rate,omitempty"`
	BitRate    int64  `json:"bit_rate,omitempty,string"`
}

type ProbeFormat struct {
	Size     int64   `json:"size,omitempty,string"`
	Duration float64 `json:"duration,omitempty,string"`
}
