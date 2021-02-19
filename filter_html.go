package subtitles

import (
	log "github.com/sirupsen/logrus"

	"github.com/kennygrant/sanitize"
)

// filterHTML removes all html tags from captions
func (subtitle *Subtitle) filterHTML() *Subtitle {
	for _, cap := range subtitle.Captions {
		for i, line := range cap.Text {
			clean := sanitize.HTML(line)
			if clean != cap.Text[i] {
				log.Println("[html]", cap.Text[i], "->", clean)
				cap.Text[i] = clean
			}
		}
	}
	return subtitle
}
