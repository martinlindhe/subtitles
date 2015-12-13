package filter

import (
	"fmt"

	"github.com/martinlindhe/go-subber/caption"
)

// FilterSubs pass the captions through a filter function
func FilterSubs(captions []caption.Caption, filter string) []caption.Caption {

	if filter == "capslock" {
		return CapslockStripper(captions)
	}
	if filter == "html" {
		return HTMLStripper(captions)
	}
	if filter != "none" {
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}

	return captions
}
