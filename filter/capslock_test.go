package filter

import (
	"fmt"
	"testing"

	"github.com/martinlindhe/go-subber/caption"
	"github.com/martinlindhe/go-subber/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieHashFromMovieFile(t *testing.T) {

	var in []caption.Caption
	in = append(in, caption.Caption{
		Seq:   1,
		Text:  []string{"GO NINJA!", "NINJA GO!"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})

	var expected []caption.Caption
	expected = append(expected, caption.Caption{
		Seq:   1,
		Text:  []string{"Go ninja!", "Ninja go!"},
		Start: common.MakeTime(0, 0, 4, 630),
		End:   common.MakeTime(0, 0, 6, 18)})

	fmt.Println(CapslockStripper(in))

	assert.Equal(t, expected, CapslockStripper(in))
}
