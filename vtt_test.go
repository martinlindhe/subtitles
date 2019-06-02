package subtitles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVTT(t *testing.T) {
	txt := "WEBVTT\n" +
		"\n" +
		"00:00:00.000 --> 00:00:05.560\n" +
		"I Vetenskapens värld: ett experiment\n" +
		"jag aldrig kommer att glömma.\n" +
		"\n" +
		"00:00:10.840 --> 00:00:15.760\n" +
		"Om en person får kämpa\n" +
		"för att hålla sig vaken–\n"

	res, err := NewFromVTT(txt)
	assert.Equal(t, nil, err)

	expected := Subtitle{
		[]Caption{{
			1,
			makeTime(0, 0, 0, 0),
			makeTime(0, 0, 5, 560),
			[]string{
				"I Vetenskapens värld: ett experiment",
				"jag aldrig kommer att glömma.",
			}}, {
			2,
			makeTime(0, 0, 10, 840),
			makeTime(0, 0, 15, 760),
			[]string{
				"Om en person får kämpa",
				"för att hålla sig vaken–",
			}}}}
	assert.Equal(t, expected, res)
}

func TestAsVTT(t *testing.T) {
	expected := "WEBVTT\n" +
		"\n" +
		"00:04.630 --> 00:06.018\n" +
		"Go ninja!\n" +
		"\n" +
		"01:09.630 --> 01:11.005\n" +
		"No ninja!\n\n"

	assert.Equal(t, true, looksLikeVTT(expected))

	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"Go ninja!"},
	}, {
		2,
		makeTime(0, 1, 9, 630),
		makeTime(0, 1, 11, 005),
		[]string{"No ninja!"},
	}}}

	assert.Equal(t, expected, in.AsVTT())
}

func ExampleNewFromSRT() {
	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"Go ninja!\n" +
		"\n" +
		"1\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"No ninja!\n"

	res, _ := NewFromSRT(in)

	// Output: WEBVTT
	//
	// 00:04.630 --> 00:06.018
	// Go ninja!
	//
	// 01:09.630 --> 01:11.005
	// No ninja!
	fmt.Println(res.AsVTT())
}
