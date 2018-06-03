package subtitles

import (
	"strings"
	"unicode"
	"unicode/utf8"

	log "github.com/Sirupsen/logrus"
)

var (
	ocrErrors = map[string]string{
		"s0 ":       "so ",
		"g0 ":       "go ",
		"0n ":       "on ",
		"c0uld":     "could",
		"s0mething": "something",
		"l've":      "I've",
		"1 Oth":     "10th",
	}
)

// filterOCR corrects some OCR mistakes
func (subtitle *Subtitle) filterOCR() *Subtitle {
	for _, cap := range subtitle.Captions {
		for i, org := range cap.Text {
			s := cap.Text[i]
			for bad, good := range ocrErrors {
				// lower case
				s = strings.Replace(s, bad, good, -1)

				// upper case
				s = strings.Replace(s, strings.ToUpper(bad), strings.ToUpper(good), -1)

				// ucfirst
				s = strings.Replace(s, strings.Title(bad), strings.Title(good), -1)
			}

			s = fixOCRLineCapitalization(s)
			if org != s {
				log.Println("[ocr]", org, "->", s)
			}
			cap.Text[i] = s
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
	if len(s) <= 3 || !isASCIIOnly(s) {
		return s
	}

	// don't touch group of lowercase + uppercase such as in "macOS"
	cases := countCaseInLetters(s)
	if len(cases) < 4 {
		return s
	}

	if countUppercaseLetters(s) >= 2 {
		return strings.ToUpper(s)
	}
	if startsWithUppercase(s) {
		return strings.Title(s)
	}
	return strings.ToLower(s)
}

func countUppercaseLetters(s string) int {
	upper := 0
	for _, c := range s {
		if unicode.IsUpper(c) {
			upper++
		}
	}
	return upper
}

func countLowercaseLetters(s string) int {
	lower := 0
	for _, c := range s {
		if unicode.IsLower(c) {
			lower++
		}
	}
	return lower
}

func startsWithUppercase(s string) bool {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return false
	}
	if size > 0 && unicode.IsUpper(r) {
		return true
	}
	return false
}

func isASCIIOnly(s string) bool {
	for _, c := range s {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}

type caseCount struct {
	kind caseType
	n    int
}

type caseType int

const (
	none caseType = iota
	lower
	upper
)

func getCase(c rune) caseType {
	if unicode.IsUpper(c) {
		return upper
	}
	if unicode.IsLower(c) {
		return lower
	}
	return none
}

func countCaseInLetters(s string) []caseCount {
	res := []caseCount{}
	currentCount := 0
	lastCase := none
	for _, c := range s {
		currentCase := getCase(c)
		if lastCase == none {
			lastCase = currentCase
		}
		if lastCase != currentCase {
			if currentCount > 0 {
				res = append(res, caseCount{lastCase, currentCount})
				currentCount = 1
				lastCase = currentCase
			}
		} else {
			currentCount++
		}
	}
	if currentCount > 0 {
		res = append(res, caseCount{lastCase, currentCount})
	}
	return res
}
