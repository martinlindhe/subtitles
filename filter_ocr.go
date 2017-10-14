package subtitles

import (
	"strings"

	log "github.com/Sirupsen/logrus"
)

var (
	ocrErrors = map[string]string{
		"s0 ":       "so ",
		"g0 ":       "go ",
		"0n ":       "on ",
		"c0uld":     "could",
		"s0mething": "something",
		"l've":      "i've",
	}
)

// filterOCR corrects some OCR mistakes
func (subtitle *Subtitle) filterOCR() *Subtitle {
	for _, cap := range subtitle.Captions {
		for i, org := range cap.Text {
			for bad, good := range ocrErrors {
				// lower case
				cap.Text[i] = strings.Replace(cap.Text[i], bad, good, -1)

				// upper case
				cap.Text[i] = strings.Replace(cap.Text[i], strings.ToUpper(bad), strings.ToUpper(good), -1)

				// ucfirst
				cap.Text[i] = strings.Replace(cap.Text[i], strings.Title(bad), strings.Title(good), -1)
			}

			if org != cap.Text[i] {
				log.Println("[ocr]", org, "->", cap.Text[i])
			}
		}
	}
	return subtitle
}
