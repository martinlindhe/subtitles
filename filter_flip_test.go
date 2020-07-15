package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFlip(t *testing.T) {
	in := Subtitle{Captions: []Caption{{
		Seq:   1,
		Start: makeTime(0, 0, 4, 630),
		End:   makeTime(0, 0, 6, 18),
		Text:  []string{"Line one", "Line two"},
	}}}
	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"Line two", "Line one"},
	}}}
	assert.Equal(t, &expected, in.filterFlip())
}
