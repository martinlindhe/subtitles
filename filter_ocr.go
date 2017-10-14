package subtitles

import (
	"strings"
	"unicode"

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

			cap.Text[i] = fixOCRLineCapitalization(cap.Text[i])
			if org != cap.Text[i] {
				log.Println("[ocr]", org, "->", cap.Text[i])
			}
		}
	}
	return subtitle
}

func fixOCRLineCapitalization(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = fixOCRWordCapitalization(words[i])
	}
	return strings.Join(words, " ")
}

// fix capitalization errors due to ocr, GAsPs => GASPS
func fixOCRWordCapitalization(s string) string {
	if len(s) <= 3 {
		return s
	}

	// if starts with uc, or at least 2 letters is upper, make all upper
	upper := 0
	ucStart := false
	for i, char := range s {
		if i == 0 && unicode.IsUpper(char) {
			ucStart = true
		}
		if unicode.IsUpper(char) {
			upper++
		}
	}
	if upper >= 2 {
		return strings.ToUpper(s)
	}
	if ucStart {
		return strings.Title(s)
	}
	return strings.ToLower(s)
}
