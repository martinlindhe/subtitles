package subtitles

import (
	"fmt"
	"time"
)

// Caption represents one subtitle block
type Caption struct {
	Seq   int
	Start time.Time
	End   time.Time
	Text  []string
}

// AsSrt renders the caption as srt
func (cap Caption) AsSrt() string {
	res := fmt.Sprintf("%d", cap.Seq) + eol +
		SrtTime(cap.Start) + " --> " + SrtTime(cap.End) + eol
	for _, line := range cap.Text {
		res += line + eol
	}
	return res + eol
}
