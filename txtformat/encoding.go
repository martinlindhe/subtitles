package txtformat

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// ConvertToUTF8 returns an utf8 string
func ConvertToUTF8(b []byte) string {

	s := ""

	if hasUTF16BeMarker(b) {
		s, _ = utf16ToUTF8(b[2:], true)
	} else if hasUTF16LeMarker(b) {
		s, _ = utf16ToUTF8(b[2:], false)
	} else if hasUTF8Marker(b) {
		s = string(b[3:])
	} else if looksLikeLatin1(b) {
		s = latin1toUTF8(b)
	} else {
		s = string(b)
	}

	return NormalizeLineFeeds(s)
}

// NormalizeLineFeeds will return a string with \n as linefeeds
func NormalizeLineFeeds(s string) string {

	if len(s) < 80 {
		return s
	}

	r := 0
	n := 0

	for i := 0; i < 80; i++ {
		if s[i] == '\r' {
			r++
		} else if s[i] == '\n' {
			n++
		}
	}

	if n == 0 && r > 0 {
		// older Mac files has \r linebreak
		return strings.Replace(s, "\r", "\n", -1)
	}

	return strings.Replace(s, "\r\n", "\n", -1)
}

func looksLikeLatin1(b []byte) bool {

	swe := float64(0)

	for i := 0; i < len(b); i++ {
		switch {
		case b[i] == 0xe5: // å
			swe++
		case b[i] == 0xe4: // ä
			swe++
		case b[i] == 0xf6: // ö
			swe++
		case b[i] == 0xc4: // Ä
			swe++
		case b[i] == 0xc5: // Å
			swe++
		case b[i] == 0xd6: // Ö
			swe++
		}
	}

	// calc percent of swe letters
	pct := (swe / float64(len(b))) * 100
	if pct >= 1 {
		return true
	}

	if pct > 0 {
		//fmt.Printf("XXX %v %% swe letters, %v\n", pct, swe)
	}

	return false
}

func latin1toUTF8(in []byte) string {

	res := make([]rune, len(in))
	for i, b := range in {
		res[i] = rune(b)
	}
	return string(res)
}

func hasUTF8Marker(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	if b[0] == 0xef && b[1] == 0xbb && b[2] == 0xbf {
		return true
	}
	return false
}
func hasUTF16BeMarker(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	if b[0] == 0xfe && b[1] == 0xff {
		return true
	}
	return false
}

func hasUTF16LeMarker(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	if b[0] == 0xff && b[1] == 0xfe {
		return true
	}
	return false
}

func utf16ToUTF8(b []byte, bigEndian bool) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		if bigEndian {
			u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		} else {
			u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		}
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
