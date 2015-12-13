package caption

import (
	"testing"

	"github.com/martinlindhe/subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestCleanSub(t *testing.T) {

	var in = []Caption{
		{
			1,
			testExtras.MakeTime(0, 0, 4, 630),
			testExtras.MakeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			testExtras.MakeTime(0, 1, 9, 630),
			testExtras.MakeTime(0, 1, 11, 005),
			[]string{"Subtitles By MrCool"},
		},
		{
			3,
			testExtras.MakeTime(0, 1, 9, 630),
			testExtras.MakeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	cleaned := CleanSubs(in)

	var expected = []Caption{
		{
			1,
			testExtras.MakeTime(0, 0, 4, 630),
			testExtras.MakeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			testExtras.MakeTime(0, 1, 9, 630),
			testExtras.MakeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	assert.Equal(t, expected, cleaned)
}
