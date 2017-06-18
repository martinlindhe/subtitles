package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterCapitalization(t *testing.T) {

	in := Subtitle{Captions: []Caption{{
		Seq:   1,
		Start: MakeTime(0, 0, 4, 630),
		End:   MakeTime(0, 0, 6, 18),
		Text:  []string{"GO NINJA!", "NINJA GO!"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"Go ninja!", "Ninja go!"},
	}}}

	assert.Equal(t, &expected, in.filterCapitalization())
}
