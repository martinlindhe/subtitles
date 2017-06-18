package subtitles

import (
	"time"
)

// Caption represents one subtitle block
type Caption struct {
	Seq   int
	Start time.Time
	End   time.Time
	Text  []string
}
