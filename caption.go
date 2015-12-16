package subber

import "time"

// caption represents one subtitle block
type caption struct {
	Seq   int
	Start time.Time
	End   time.Time
	Text  []string
}
