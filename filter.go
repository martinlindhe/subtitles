package subtitles

import (
	"fmt"
)

// FilterCaptions pass the captions through a filter function
func (subtitle *Subtitle) FilterCaptions(filter string) {
	switch filter {
	case "all":
		subtitle.filterCapitalization()
		subtitle.filterHTML()
		subtitle.filterOCR()
	case "caps":
		subtitle.filterCapitalization()
	case "html":
		subtitle.filterHTML()
	case "ocr":
		subtitle.filterOCR()
	case "flip":
		subtitle.filterFlip()
	case "none":
	default:
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}
}
