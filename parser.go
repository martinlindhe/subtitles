package subtitles

// parse tries to parse a subtitle from the data stream
func parse(b []byte) (Subtitle, error) {

	s := convertToUTF8(b)

	if s[0] == '[' {
		// looks like ssa
		return NewFromSSA(s)
	}

	// XXX
	return NewFromSRT(s)
}
