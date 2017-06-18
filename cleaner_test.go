package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAds(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"Go ninja!"},
	}, {
		2,
		MakeTime(0, 1, 9, 630),
		MakeTime(0, 1, 11, 005),
		[]string{"Subtitles By MrCool"},
	}, {
		3,
		MakeTime(0, 1, 9, 630),
		MakeTime(0, 1, 11, 005),
		[]string{"No ninja!"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"Go ninja!"},
	}, {
		2,
		MakeTime(0, 1, 9, 630),
		MakeTime(0, 1, 11, 005),
		[]string{"No ninja!"},
	}}}

	assert.Equal(t, &expected, in.removeAds())
}
