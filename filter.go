package subtitles

import (
	"fmt"
)

// filterSubs pass the captions through a filter function
func (subtitle *Subtitle) filterSubs(filter string) {
	if filter == "caps" {
		subtitle.filterCapitalization()
	}
	if filter == "html" {
		subtitle.filterHTML()
	}
	if filter != "none" {
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}
}
