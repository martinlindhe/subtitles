package subtitles

import (
	"time"
)

// AsVTT renders the sub in WebVTT format
// https://en.wikipedia.org/wiki/WebVTT
func (subtitle *Subtitle) AsVTT() (res string) {
	res = "WEBVTT\n\n"
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
	if t.Hour() == 0 {
		return t.Format("04:05.000")
	}
	return t.Format("15:04:05.000")
}
