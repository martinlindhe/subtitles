package subtitles

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

// filterCapitalization converts "ALL CAPS" text into "Initial letter capped"
func (subtitle *Subtitle) filterCapitalization() *Subtitle {
	for _, cap := range subtitle.Captions {
		for i, line := range cap.Text {

			clean := ucFirst(line)
			if clean != cap.Text[i] {
				log.Println("[caps]", cap.Text[i], "-->", clean)
				cap.Text[i] = clean
			}
		}
	}
	return subtitle
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
