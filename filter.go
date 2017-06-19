package subtitles

import (
	"fmt"
)

// FilterCaptions pass the captions through a filter function
func (subtitle *Subtitle) FilterCaptions(filter string) {
	switch filter {
	case "caps":
		subtitle.filterCapitalization()
	case "html":
		subtitle.filterHTML()
	case "none":
	default:
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}
}
