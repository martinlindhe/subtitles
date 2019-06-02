package subtitles

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var webVTTTag = "WEBVTT\n"

// AsVTT renders the sub in WebVTT format
// https://en.wikipedia.org/wiki/WebVTT
func (subtitle *Subtitle) AsVTT() (res string) {
	res = webVTTTag + "\n"
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

func looksLikeVTT(s string) bool {
	return strings.HasPrefix(s, webVTTTag)
}

// NewFromVTT parses a .vtt text into Subtitle, assumes s is a clean utf8 string
func NewFromVTT(s string) (res Subtitle, err error) {
	idx := strings.Index(s, webVTTTag)
	if idx == -1 {
		return res, fmt.Errorf("not a vtt")
	}
	s = s[idx+len(webVTTTag):]

	r1 := regexp.MustCompile("([0-9:.,]*) --> ([0-9:.,]*)")
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i := 0; i < len(lines); i++ {
		seq := strings.Trim(lines[i], "\r ")
		if seq == "" {
			continue
		}

		var o Caption
		o.Seq = outSeq

		matches := r1.FindStringSubmatch(lines[i])
		if len(matches) < 3 {
			err = fmt.Errorf("vtt: parse error at line %d (idx out of range) for input '%s'", i, lines[i])
			break
		}

		o.Start, err = parseTime(matches[1])
		if err != nil {
			err = fmt.Errorf("vtt: start error at line %d: %v", i, err)
			break
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			err = fmt.Errorf("vtt: end error at line %d: %v", i, err)
			break
		}

		i++
		if i >= len(lines) {
			break
		}

		textLine := 1
		for {
			line := strings.Trim(lines[i], "\r ")
			if line == "" && textLine > 1 {
				break
			}
			if line != "" {
				o.Text = append(o.Text, line)
			}
			i++
			if i >= len(lines) {
				break
			}
			textLine++
		}

		if len(o.Text) > 0 {
			res.Captions = append(res.Captions, o)
			outSeq++
		}
	}
	return
}
