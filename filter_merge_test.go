package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMerge(t *testing.T) {
	in := `1
00:01:28,800 --> 00:01:30,166
I never thought

2
00:01:28,800 --> 00:01:30,166
I would be caught up

3
00:01:30,266 --> 00:01:32,233
in a story such as this,

4
00:01:33,266 --> 00:01:35,867
because I live on the

5
00:01:33,266 --> 00:01:35,867
other side of the world.
`

	sub, err := NewFromSRT(in)
	assert.Equal(t, nil, err)

	sub = *sub.filterMerge()

	expected := `1
00:01:28,800 --> 00:01:30,166
I never thought
I would be caught up

3
00:01:30,266 --> 00:01:32,233
in a story such as this,

4
00:01:33,266 --> 00:01:35,867
because I live on the
other side of the world.

`
	assert.Equal(t, expected, sub.AsSRT())

}
