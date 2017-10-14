package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterOCRLower(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"s0mething good"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"something good"},
	}}}

	assert.Equal(t, &expected, in.filterOCR())
}

func TestFilterOCRUpper(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"S0METHING GOOD"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"SOMETHING GOOD"},
	}}}

	assert.Equal(t, &expected, in.filterOCR())
}

func TestFilterOCRUcFirst(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"S0mething good"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"Something good"},
	}}}

	assert.Equal(t, &expected, in.filterOCR())
}

func TestFilterOCREnglish(t *testing.T) {

	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"l've got a feeling"},
	}}}

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"i've got a feeling"},
	}}}

	assert.Equal(t, &expected, in.filterOCR())
}
