package caption

import (
	"testing"

	"github.com/martinlindhe/go-subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestRenderTime(t *testing.T) {

	cap := Caption{
		1,
		testExtras.MakeTime(18, 40, 22, 110),
		testExtras.MakeTime(18, 41, 20, 123),
		[]string{"<i>Go ninja!</i>"},
	}

	assert.Equal(t, "18:40:22,110 --> 18:41:20,123", cap.SrtTime())
}
