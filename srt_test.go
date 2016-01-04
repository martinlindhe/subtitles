package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {

	t1, _ := parseTime("18:40:22.110")
	t2, _ := parseTime("18:40:22,110")
	t3, _ := parseTime("18:40:22")
	t4, _ := parseTime("00:00:0,500")
	t5, _ := parseTime("00:00:2,00")
	t6, _ := parseTime("00:14:52.12")

	assert.Equal(t, makeTime(18, 40, 22, 110), t1)
	assert.Equal(t, makeTime(18, 40, 22, 110), t2)
	assert.Equal(t, makeTime(18, 40, 22, 0), t3)
	assert.Equal(t, makeTime(0, 0, 0, 500), t4)
	assert.Equal(t, makeTime(0, 0, 2, 0), t5)
	assert.Equal(t, makeTime(0, 14, 52, 12), t6)
}

func TestParseSrt(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"Go ninja!\n" +
		"\n" +
		"2\n" +
		"00:00:10,000 --> 00:00:11,000\n" +
		"Subtitles By MrCool\n" +
		"\n" +
		"3\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"No ninja!\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			makeTime(0, 0, 10, 0),
			makeTime(0, 0, 11, 0),
			[]string{"Subtitles By MrCool"},
		},
		{
			3,
			makeTime(0, 1, 9, 630),
			makeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	assert.Equal(t, expected, parseSrt(in))
}

func TestParseSrtWithMacLinebreaks(t *testing.T) {

	in := "1\r" +
		"00:00:04,630 --> 00:00:06,018\r" +
		"Go ninja!\r" +
		"\r" +
		"3\r" +
		"00:01:09,630 --> 00:01:11,005\r" +
		"No ninja!\r"

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			makeTime(0, 1, 9, 630),
			makeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	utf8 := convertToUTF8([]byte(in))

	assert.Equal(t, expected, parseSrt(utf8))
}

func TestParseSrtSkipEmpty(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"Go ninja!\n" +
		"\n" +
		"2\n" +
		"00:00:10,000 --> 00:00:11,000\n" +
		"\n" +
		"\n" +
		"3\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"No ninja!\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			makeTime(0, 1, 9, 630),
			makeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	assert.Equal(t, expected, parseSrt(in))
}

func TestParseSrtCrlf(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		"Go ninja!\r\n" +
		"\r\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
	}

	assert.Equal(t, expected, parseSrt(in))
}

func TestParseExtraLineBreak(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		// NOTE: should not be line break here, but some files has,
		// so lets make sure we handle it
		"\r\n" +
		"Go ninja!\r\n" +
		"\r\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
	}

	assert.Equal(t, expected, parseSrt(in))
}

func TestParseWierdTimestamp(t *testing.T) {

	in := "1\r\n" +
		"00:14:52.00 --> 00:14:57,500\r\n" +
		"Go ninja!\r\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 14, 52, 0),
			makeTime(0, 14, 57, 500),
			[]string{"Go ninja!"},
		},
	}

	assert.Equal(t, expected, parseSrt(in))
}

func TestRenderSrt(t *testing.T) {

	expected := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"Go ninja!\n" +
		"\n" +
		"2\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"No ninja!\n\n"

	var in = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"Go ninja!"},
		},
		{
			2,
			makeTime(0, 1, 9, 630),
			makeTime(0, 1, 11, 005),
			[]string{"No ninja!"},
		},
	}

	assert.Equal(t, expected, renderSrt(in))
}

func TestParseLatin1Srt(t *testing.T) {
	in := "1\r\n" +
		"00:14:52.00 --> 00:14:57,500\r\n" +
		"Hall\xe5 ninja!\r\n"

	var expected = []caption{
		{
			1,
			makeTime(0, 14, 52, 0),
			makeTime(0, 14, 57, 500),
			[]string{"Hallå ninja!"},
		},
	}

	utf8 := convertToUTF8([]byte(in))

	assert.Equal(t, expected, parseSrt(utf8))
}

func TestParseUTF16BESrt(t *testing.T) {

	in := []byte{
		0xfe, 0xff, // UTF16 BE BOM

		0, '1',
		0, '\r', 0, '\n',

		0, '0', 0, '0', 0, ':', 0, '0', 0, '0', 0, ':',
		0, '0', 0, '0', 0, ',', 0, '0', 0, '0', 0, '0',

		0, ' ', 0, '-', 0, '-', 0, '>', 0, ' ',

		0, '0', 0, '0', 0, ':', 0, '0', 0, '0', 0, ':',
		0, '0', 0, '0', 0, ',', 0, '0', 0, '0', 0, '1',

		0, '\r', 0, '\n',

		0, 'T', 0, 'e', 0, 's', 0, 't',

		0, '\r', 0, '\n',
		0, '\r', 0, '\n',
	}

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 0, 0),
			makeTime(0, 0, 0, 1),
			[]string{"Test"},
		},
	}

	utf8 := convertToUTF8(in)

	assert.Equal(t, expected, parseSrt(utf8))
}

func TestParseUTF16LESrt(t *testing.T) {

	in := []byte{
		0xff, 0xfe, // UTF16 LE BOM

		'1', 0,
		'\r', 0, '\n', 0,

		'0', 0, '0', 0, ':', 0, '0', 0, '0', 0, ':', 0,
		'0', 0, '0', 0, ',', 0, '0', 0, '0', 0, '0', 0,

		' ', 0, '-', 0, '-', 0, '>', 0, ' ', 0,

		'0', 0, '0', 0, ':', 0, '0', 0, '0', 0, ':', 0,
		'0', 0, '0', 0, ',', 0, '0', 0, '0', 0, '1', 0,

		'\r', 0, '\n', 0,

		'T', 0, 'e', 0, 's', 0, 't', 0,

		'\r', 0, '\n', 0,
		'\r', 0, '\n', 0,
	}

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 0, 0),
			makeTime(0, 0, 0, 1),
			[]string{"Test"},
		},
	}

	utf8 := convertToUTF8(in)

	assert.Equal(t, expected, parseSrt(utf8))
}

func TestParseUTF8BomSrt(t *testing.T) {

	in := []byte{
		0xef, 0xbb, 0xbf, // UTF8 BOM

		'1',
		'\r', '\n',

		'0', '0', ':', '0', '0', ':',
		'0', '0', ',', '0', '0', '0',

		' ', '-', '-', '>', ' ',

		'0', '0', ':', '0', '0', ':',
		'0', '0', ',', '0', '0', '1',

		'\r', '\n',

		'T', 'e', 's', 't',

		'\r', '\n',
		'\r', '\n',
	}

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 0, 0),
			makeTime(0, 0, 0, 1),
			[]string{"Test"},
		},
	}

	utf8 := convertToUTF8(in)

	assert.Equal(t, expected, parseSrt(utf8))
}
