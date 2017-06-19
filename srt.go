package subtitles

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Eol is the end of line characters to use when writing .srt data
const eol = "\n"

func looksLikeSRT(s string) bool {
	if strings.HasPrefix(s, "1\n") || strings.HasPrefix(s, "1\r\n") {
		return true
	}
	return false
}

// NewFromSRT parses a .srt text into Subtitle, assumes s is a clean utf8 string
func NewFromSRT(s string) (res Subtitle, err error) {
	r1 := regexp.MustCompile("([0-9:.,]*) --> ([0-9:.,]*)")
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i := 0; i < len(lines); i++ {
		seq := strings.Trim(lines[i], "\r ")
		if seq == "" {
			continue
		}

		_, err := strconv.Atoi(seq)
		if err != nil {
			err = fmt.Errorf("srt: atoi error at line %d: %v", i, err)
			break
		}

		var o Caption
		o.Seq = outSeq

		i++
		if i >= len(lines) {
			break
		}

		matches := r1.FindStringSubmatch(lines[i])
		if len(matches) < 3 {
			err = fmt.Errorf("srt: parse error at line %d (idx out of range)", i)
			break
		}

		o.Start, err = parseTime(matches[1])
		if err != nil {
			err = fmt.Errorf("srt: start error at line %d: %v", i, err)
			break
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			err = fmt.Errorf("srt: end error at line %d: %v", i, err)
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

// AsSRT renders the sub in .srt format
func (subtitle *Subtitle) AsSRT() (res string) {
	for _, sub := range subtitle.Captions {
		res += sub.AsSRT()
	}
	return
}

// AsSRT renders the caption as srt
func (cap Caption) AsSRT() string {
	res := fmt.Sprintf("%d", cap.Seq) + eol +
		TimeSRT(cap.Start) + " --> " + TimeSRT(cap.End) + eol
	for _, line := range cap.Text {
		res += line + eol
	}
	return res + eol
}

// TimeSRT renders a timestamp for use in .srt
func TimeSRT(t time.Time) string {
	res := t.Format("15:04:05.000")
	return strings.Replace(res, ".", ",", 1)
}
