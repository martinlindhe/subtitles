package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterCapitalization(t *testing.T) {

	in := Subtitle{Captions: []Caption{{
		Seq:   1,
		Start: makeTime(0, 0, 4, 630),
		End:   makeTime(0, 0, 6, 18),
		Text:  []string{"GO NINJA!", "NINJA GO!"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"Go ninja!", "Ninja go!"},
	}}}

	assert.Equal(t, &expected, in.filterCapitalization())
}
