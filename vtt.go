package subtitles

import (
	"time"
)

// AsVTT renders the sub in WebVTT format
// https://en.wikipedia.org/wiki/WebVTT
func (subtitle *Subtitle) AsVTT() (res string) {
	res = "WEBVTT\n\n" // XXX
	for _, sub := range subtitle.Captions {
		res += sub.AsVTT()
	}
	return
}

// AsVTT renders the caption as WebVTT
func (cap Caption) AsVTT() string {
	res := TimeVTT(cap.Start) + " --> " + TimeVTT(cap.End) + eol
	for _, line := range cap.Text {
		res += line + eol
	}
	return res + eol
}

// TimeVTT renders a timestamp for use in WebVTT
func TimeVTT(t time.Time) string {
	// XXX hours are optional, size-optimize!
	return t.Format("15:04:05.000")
}
