package subtitles

import (
	"log"

	"github.com/kennygrant/sanitize"
)

// filterHTML removes all html tags from captions
func filterHTML(captions []caption) []caption {

	for _, cap := range captions {
		for i, line := range cap.Text {
			clean := sanitize.HTML(line)

			if clean != cap.Text[i] {
				log.Printf("[html] %s -> %s\n", cap.Text[i], clean)
				cap.Text[i] = clean
			}
		}
	}

	return captions
}
