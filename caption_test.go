package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTime(t *testing.T) {

	cap := caption{
		1,
		makeTime(18, 40, 22, 110),
		makeTime(18, 41, 20, 123),
		[]string{"<i>Go ninja!</i>"},
	}

	assert.Equal(t, "18:40:22,110 --> 18:41:20,123", cap.srtTime())
}
