package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterHTML(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"<b>GO NINJA!</b>", "NINJA&nbsp;GO!"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"GO NINJA!", "NINJA GO!"},
	}}}

	assert.Equal(t, &expected, in.filterHTML())
}
