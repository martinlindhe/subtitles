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
		[]string{"I've got a feeling"},
	}}}
	assert.Equal(t, &expected, in.filterOCR())
}

func TestFilterOCRCapitalization(t *testing.T) {
	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"GAsPs slowly"},
	}}}
	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"GASPS slowly"},
	}}}
	assert.Equal(t, &expected, in.filterOCR())
}

func TestFixOCRWordCapitalization(t *testing.T) {
	assert.Equal(t, "He's", fixOCRWordCapitalization("He's"))
	assert.Equal(t, "GASPS", fixOCRWordCapitalization("GAsPs"))

	assert.Equal(t, "macOS", fixOCRWordCapitalization("macOS"))
	assert.Equal(t, "WindowsXP", fixOCRWordCapitalization("WindowsXP"))
}

func TestStartsWithUppercase(t *testing.T) {
	assert.Equal(t, true, startsWithUppercase("Allo"))
	assert.Equal(t, true, startsWithUppercase("Ällo"))
	assert.Equal(t, false, startsWithUppercase("allo"))
}

func TestCountCaseInLetters(t *testing.T) {
	assert.Equal(t, []caseCount{{upper, 2}}, countCaseInLetters("GA"))
	assert.Equal(t, []caseCount{{lower, 2}}, countCaseInLetters("ga"))
	assert.Equal(t, []caseCount{{upper, 2}, {lower, 1}, {upper, 1}, {lower, 1}}, countCaseInLetters("GAsPs"))
}
