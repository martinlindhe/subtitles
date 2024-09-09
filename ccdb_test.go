package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromCCDBCapture(t *testing.T) {

	in := "[SUBTITLE]\r\n" +
		"[COLF]&HFFFFFF,[STYLE]no,[SIZE]10,[FONT]Arial\r\n" +
		"00:00:16.24,00:01:25.82\r\n" +
		"Whoa.                           \r\n" +
		"\r\n" +
		"00:01:31.45,00:01:33.62\r\n" +
		"Go on. Get out.                 \r\n" +
		"\r\n" +
		"00:01:33.62,00:01:33.65\r\n" +
		"                                \r\n" + // should disappear in the parsed captions
		"\r\n" +
		"00:01:33.65,00:01:34.81\r\n" +
		"Out!                            \r\n" +
		"\r\n"

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 16, 24),
		makeTime(0, 1, 25, 82),
		[]string{"Whoa."},
	}, {
		2,
		makeTime(0, 1, 31, 45),
		makeTime(0, 1, 33, 62),
		[]string{"Go on. Get out."},
	}, {
		3,
		makeTime(0, 1, 33, 65),
		makeTime(0, 1, 34, 81),
		[]string{"Out!"},
	}}}

	res, err := NewFromCCDBCapture(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}

func TestParseCCDBTime(t *testing.T) {
	t1, _ := parseCCDBTime("00:00:16.24")
	assert.Equal(t, makeTime(0, 0, 16, 24), t1)
}
