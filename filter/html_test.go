package filter

import (
	"testing"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestHTMLStripper(t *testing.T) {

	var in = []caption.Caption{
		{
			1,
			testExtras.MakeTime(0, 0, 4, 630),
			testExtras.MakeTime(0, 0, 6, 18),
			[]string{"<b>GO NINJA!</b>", "NINJA&nbsp;GO!"},
		},
	}

	var expected = []caption.Caption{
		{
			1,
			testExtras.MakeTime(0, 0, 4, 630),
			testExtras.MakeTime(0, 0, 6, 18),
			[]string{"GO NINJA!", "NINJA GO!"},
		},
	}

	assert.Equal(t, expected, HTMLStripper(in))
}
