package subtitles

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromDCSubGood(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/dcsub/good.xml")
	assert.Equal(t, nil, err)

	in := string(data)
	expected := Subtitle{[]Caption{{
		Seq:   1,
		Start: makeTime(0, 1, 0, 0),
		End:   makeTime(0, 1, 3, 240),
		Text:  []string{"-Nej...", "-Vad är det?"},
	}}}

	res, err := NewFromDCSub(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}

func TestNewFromDCSubStyledText(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/dcsub/styled_text.xml")
	assert.Equal(t, nil, err)

	in := string(data)
	expected := Subtitle{[]Caption{{
		Seq:   1,
		Start: makeTime(0, 5, 40, 0),
		End:   makeTime(0, 5, 43, 0),
		Text:  []string{`<font italic="yes">Hanne!</font>`},
	}}}

	res, err := NewFromDCSub(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}

func TestNewFromDCSubEntitiesText(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/dcsub/entities.xml")
	assert.Equal(t, nil, err)

	in := string(data)
	expected := Subtitle{[]Caption{{
		Seq:   1,
		Start: makeTime(0, 1, 0, 0),
		End:   makeTime(0, 1, 3, 240),
		Text:  []string{`Eller så är det en "ähpar".`},
	}}}

	res, err := NewFromDCSub(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}
