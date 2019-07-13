package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSrtTime(t *testing.T) {
	t1, _ := parseSrtTime("18:40:22.110")
	t2, _ := parseSrtTime("18:40:22,110")
	t3, _ := parseSrtTime("18:40:22:110")
	t4, _ := parseSrtTime("18:40:22")
	t5, _ := parseSrtTime("00:00:0,500")
	t6, _ := parseSrtTime("00:00:2,00")
	t7, _ := parseSrtTime("00:14:52.12")

	assert.Equal(t, makeTime(18, 40, 22, 110), t1)
	assert.Equal(t, makeTime(18, 40, 22, 110), t2)
	assert.Equal(t, makeTime(18, 40, 22, 110), t3)
	assert.Equal(t, makeTime(18, 40, 22, 0), t4)
	assert.Equal(t, makeTime(0, 0, 0, 500), t5)
	assert.Equal(t, makeTime(0, 0, 2, 0), t6)
	assert.Equal(t, makeTime(0, 14, 52, 12), t7)
}

func TestParseVttTime(t *testing.T) {
	t1, _ := parseVttTime("00:00:10.840")
	t2, _ := parseVttTime("00:13.000")

	assert.Equal(t, makeTime(0, 0, 13, 0), t2)
	assert.Equal(t, makeTime(0, 0, 10, 840), t1)
}
