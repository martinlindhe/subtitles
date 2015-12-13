package caption

import (
	"testing"

	"github.com/martinlindhe/go-subber/common"
	"github.com/stretchr/testify/assert"
)

func TestCleanSub(t *testing.T) {

	var in []Caption
	in = append(in, Caption{
		Seq:   1,
		Text:  []string{"Go ninja!"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})
	in = append(in, Caption{
		Seq:   2,
		Text:  []string{"Subtitles By MrCool"},
		Start: common.MakeTime(0, 1, 9, 630),
		End:   common.MakeTime(0, 1, 11, 005)})
	in = append(in, Caption{
		Seq:   3,
		Text:  []string{"No ninja!"},
		Start: common.MakeTime(0, 1, 9, 630),
		End:   common.MakeTime(0, 1, 11, 005)})

	cleaned := CleanSubs(in)

	var expected []Caption
	expected = append(expected, Caption{
		Seq:   1,
		Text:  []string{"Go ninja!"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})
	expected = append(expected, Caption{
		Seq:   2,
		Text:  []string{"No ninja!"},
		Start: common.MakeTime(0, 1, 9, 630),
		End:   common.MakeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, cleaned)
}
