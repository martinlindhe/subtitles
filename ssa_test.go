package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromSsa(t *testing.T) {

	in := "[Events]\n" +
		"Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text\n" +
		"Dialogue: 0,0:01:06.37,0:01:08.04,Default,,0000,0000,0000,,Honey, I'm home!\n" +
		"Dialogue: 0,0:01:09.05,0:01:10.69,Default,,0000,0000,0000,,Hi.\\n- Hi, love.\n"

	expected := Subtitle{[]Caption{{
		1,
		MakeTime(0, 1, 6, 370),
		MakeTime(0, 1, 8, 40),
		[]string{"Honey, I'm home!"},
	}, {
		2,
		MakeTime(0, 1, 9, 50),
		MakeTime(0, 1, 10, 690),
		[]string{"Hi.", "- Hi, love."},
	}}}

	res, err := NewFromSSA(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}

func TestParseSsaFormat(t *testing.T) {

	assert.Equal(t, -1, parseSsaFormat("xxx", "some"))

	assert.Equal(t, 9, parseSsaFormat("Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text", "Text"))
}

func TestParseSsaTime(t *testing.T) {

	t1, _ := parseSsaTime("0:01:06.37")
	assert.Equal(t, MakeTime(0, 1, 6, 370), t1)
}

func TestColumnCountFromSsaFormat(t *testing.T) {
	columns := columnCountFromSsaFormat("Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text")

	assert.Equal(t, 10, columns)
}
