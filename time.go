package subtitles

import "time"

// MakeTime is a helper to create a subtitle duration
func MakeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
}
