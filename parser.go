package subtitles

// parse tries to parse a subtitle from the data stream
func parse(b []byte) []caption {

	s := convertToUTF8(b)

	if s[0] == '[' {
		// looks like ssa
		return parseSsa(s)
	}

	// XXXX
	return parseSrt(s)
}
