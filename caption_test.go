package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTime(t *testing.T) {

	cap := Caption{
		1,
		MakeTime(18, 40, 22, 110),
		MakeTime(18, 41, 20, 123),
		[]string{"<i>Go ninja!</i>"},
	}

	assert.Equal(t, "1\n"+
		"18:40:22,110 --> 18:41:20,123\n"+
		"<i>Go ninja!</i>\n\n",
		cap.AsSRT())
}
