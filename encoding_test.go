package subtitles

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLooksLikeLatin1(t *testing.T) {

	assert.Equal(t, true, looksLikeLatin1([]byte("hall\xe5")))
	assert.Equal(t, false, looksLikeLatin1([]byte("hallå")))
}

func TestReadFileAsUTF8(t *testing.T) {

	f, err := os.Open("README.md")
	assert.Equal(t, nil, err)

	_, err = readAsUTF8(f)
	assert.Equal(t, nil, err)
}
