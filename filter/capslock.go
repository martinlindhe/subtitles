package filter

import (
	"strings"

	"github.com/martinlindhe/go-subber/caption"
)

// CapslockStripper converts "ALL CAPS" text into "Initial letter capped"
func CapslockStripper(captions []caption.Caption) []caption.Caption {

	for _, cap := range captions {
		for i, line := range cap.Text {
			cap.Text[i] = ucFirst(line)
		}
	}

	return captions
}

func ucFirst(s string) string {

	res := ""

	for i, c := range s {
		if i == 0 {
			res += strings.ToUpper(string(c))
		} else {
			res += strings.ToLower(string(c))
		}
	}

	return res
}
