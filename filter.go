package subber

import (
	"fmt"
)

// filterSubs pass the captions through a filter function
func filterSubs(captions []caption, filter string) []caption {

	if filter == "caps" {
		return filterCapitalization(captions)
	}
	if filter == "html" {
		return filterHTML(captions)
	}
	if filter != "none" {
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}

	return captions
}
