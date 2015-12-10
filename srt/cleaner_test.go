package srt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanSub(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"<i>Go ninja!</i>\n" +
		"\n" +
		"2\n" +
		"00:00:10,000 --> 00:00:11,000\n" +
		"<i>Subtitles By MrCool</i>\n" +
		"\n" +
		"3\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"<i>No ninja!</i>\n"

	cleaned := CleanSubs(ParseSrt(in))

	var expected []Caption
	expected = append(expected, Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(0, 0, 4, 630), end: makeTime(0, 0, 6, 18)})
	expected = append(expected, Caption{seq: 2, text: []string{"<i>No ninja!</i>"}, start: makeTime(0, 1, 9, 630), end: makeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, cleaned)
}
