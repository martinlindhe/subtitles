package filter

import (
	"testing"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestCapsFixer(t *testing.T) {

	var in = []caption.Caption{
		{
			Seq:   1,
			Start: testExtras.MakeTime(0, 0, 4, 630),
			End:   testExtras.MakeTime(0, 0, 6, 18),
			Text:  []string{"GO NINJA!", "NINJA GO!"},
		},
	}

	var expected = []caption.Caption{
		{
			1,
			testExtras.MakeTime(0, 0, 4, 630),
			testExtras.MakeTime(0, 0, 6, 18),
			[]string{"Go ninja!", "Ninja go!"},
		},
	}

	assert.Equal(t, expected, CapsFixer(in))
}
