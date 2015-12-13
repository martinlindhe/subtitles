package filter

import (
	"fmt"
	"strings"

	"github.com/martinlindhe/go-subber/caption"
)

// CapslockStripper converts "ALL CAPS" text into "Initial letter capped"
func CapslockStripper(captions []caption.Caption) []caption.Caption {

	for _, cap := range captions {
		for i, line := range cap.Text {

			clean := ucFirst(line)

			if clean != cap.Text[i] {
				fmt.Printf("[capslock] %s -> %s\n", cap.Text[i], clean)
				cap.Text[i] = clean
			}
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
