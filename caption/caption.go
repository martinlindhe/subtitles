package caption

import (
	"strings"
	"time"
)

// Caption represents one subtitle block
type Caption struct {
	Seq   int
	Start time.Time
	End   time.Time
	Text  []string
}

func renderSrtTime(t time.Time) string {
	res := t.Format("15:04:05.000")
	return strings.Replace(res, ".", ",", 1)
}

func (cap Caption) SrtTime() string {
	return renderSrtTime(cap.Start) + " --> " + renderSrtTime(cap.End)
}
