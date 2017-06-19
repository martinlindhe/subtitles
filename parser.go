package subtitles

import "fmt"

// Parse tries to parse a subtitle
func Parse(b []byte) (Subtitle, error) {
	s := ConvertToUTF8(b)
	if looksLikeSSA(s) {
		return NewFromSSA(s)
	} else if looksLikeDCSub(s) {
		return NewFromDCSub(s)
	} else if looksLikeSRT(s) {
		return NewFromSRT(s)
	}
	return Subtitle{}, fmt.Errorf("parse: unrecognized subtitle type")
}
