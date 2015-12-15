package parser

import (
	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/srt"
	"github.com/martinlindhe/subber/ssa"
	"github.com/martinlindhe/subber/txtformat"
)

// Parse tries to parse a subtitle from the data stream
func Parse(b []byte) []caption.Caption {

	s := txtformat.ConvertToUTF8(b)

	if s[0] == '[' {
		// looks like ssa
		return ssa.ParseSsa(s)
	}

	// XXXX
	return srt.ParseSrt(s)
}
