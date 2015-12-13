package filter

import (
	"fmt"

	"github.com/kennygrant/sanitize"
	"github.com/martinlindhe/subber/caption"
)

// HTMLStripper removes all html tags from captions
func HTMLStripper(captions []caption.Caption) []caption.Caption {

	for _, cap := range captions {
		for i, line := range cap.Text {
			clean := sanitize.HTML(line)

			if clean != cap.Text[i] {
				fmt.Printf("[html] %s -> %s\n", cap.Text[i], clean)
				cap.Text[i] = clean
			}
		}
	}

	return captions
}
