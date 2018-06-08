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
	input := map[string]string{
		"GAsPs":     "GASPS",
		"He's":      "He's",
		"macOS":     "macOS",
		"WindowsXP": "WindowsXP",
	}
	for in, out := range input {
		assert.Equal(t, out, fixOCRWordCapitalization(in))
	}
}

func TestStartsWithUppercase(t *testing.T) {
	input := map[string]bool{
		"Allo": true,
		"Ällo": true,
		"allo": false,
	}
	for in, out := range input {
		assert.Equal(t, out, startsWithUppercase(in))
	}
}

func TestCountCaseInLetters(t *testing.T) {
	input := map[string][]caseCount{
		"HELLO": []caseCount{{upper, 5}},
		"hello": []caseCount{{lower, 5}},
		"Hello": []caseCount{{upper, 1}, {lower, 4}},
		"GAsPs": []caseCount{{upper, 2}, {lower, 1}, {upper, 1}, {lower, 1}},
	}
	for in, out := range input {
		assert.Equal(t, out, countCaseInLetters(in))
	}
}
