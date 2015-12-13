package srt

import (
	"testing"

	"github.com/martinlindhe/go-subber/caption"
	"github.com/martinlindhe/go-subber/common"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {

	t1, _ := parseTime("18:40:22.110")
	t2, _ := parseTime("18:40:22,110")
	t3, _ := parseTime("18:40:22")

	assert.Equal(t, common.MakeTime(18, 40, 22, 110), t1)
	assert.Equal(t, common.MakeTime(18, 40, 22, 110), t2)
	assert.Equal(t, common.MakeTime(18, 40, 22, 0), t3)
}

func TestParseSrt(t *testing.T) {

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

	var expected []caption.Caption
	expected = append(expected, caption.Caption{
		Seq:   1,
		Text:  []string{"<i>Go ninja!</i>"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})
	expected = append(expected,
		caption.Caption{
			Seq:   2,
			Text:  []string{"<i>Subtitles By MrCool</i>"},
			Start: common.MakeTime(0, 0, 10, 0),
			End:   common.MakeTime(0, 0, 11, 0)})
	expected = append(expected, caption.Caption{
		Seq:   3,
		Text:  []string{"<i>No ninja!</i>"},
		Start: common.MakeTime(0, 1, 9, 630),
		End:   common.MakeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestParseSrtCrlf(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		"<i>Go ninja!</i>\r\n" +
		"\r\n"

	var expected []caption.Caption
	expected = append(expected, caption.Caption{
		Seq:   1,
		Text:  []string{"<i>Go ninja!</i>"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestParseSrtUtf8Bom(t *testing.T) {

	in := "\ufeff1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		"<i>Go ninja!</i>\r\n" +
		"\r\n"

	var expected []caption.Caption
	expected = append(expected, caption.Caption{
		Seq:   1,
		Text:  []string{"<i>Go ninja!</i>"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestRenderSrt(t *testing.T) {

	expected := "1" + Eol +
		"00:00:04,630 --> 00:00:06,018" + Eol +
		"<i>Go ninja!</i>" + Eol +
		Eol +
		"2" + Eol +
		"00:01:09,630 --> 00:01:11,005" + Eol +
		"<i>No ninja!</i>" + Eol + Eol

	var in []caption.Caption
	in = append(in, caption.Caption{
		Seq:   1,
		Text:  []string{"<i>Go ninja!</i>"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})
	in = append(in, caption.Caption{
		Seq:   2,
		Text:  []string{"<i>No ninja!</i>"},
		Start: common.MakeTime(0, 1, 9, 630),
		End:   common.MakeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, RenderSrt(in))
}

/*
func TestWriteSrt(t *testing.T) {
	xxx
}*/
