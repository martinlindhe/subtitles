package subtitles

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFromTheSubDb(t *testing.T) {

	fileName := createZeroedTempFile(1024 * 1024 * 4)
	defer os.Remove(fileName)

	f, err := os.Open(fileName)
	assert.Equal(t, nil, err)

	hash, err := SubDbHashFromFile(f)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0dfbe8aa4c20b52e1b8bf3cb6cbdf193", hash)

	finder := NewSubFinder(f, fileName, "en")

	text, err := finder.TheSubDb("sandbox.thesubdb.com")
	assert.Equal(t, nil, err)
	assert.True(t, len(text) > 1000)
}

func subDbConformTest(t *testing.T, fileName string, expectedHash string) {
	if !exists(fileName) {
		fmt.Println("ERROR thesubdb.com conformance tests missing, run ./hash-conformance-deps if you want to run these tests")
		return
	}

	f, err := os.Open(fileName)
	assert.Equal(t, nil, err)

	hash, err := SubDbHashFromFile(f)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedHash, hash)
}

func TestSubDbHashFromFile(t *testing.T) {

	// NOTE for this to work, run "./hash-conformance-deps" to fetch needed files

	// http://thesubdb.com/api/samples/dexter.mp4
	subDbConformTest(t, "conformance-files/thesubdb/dexter.mp4", "ffd8d4aa68033dc03d1c8ef373b9028c")

	// http://thesubdb.com/api/samples/justified.mp4
	subDbConformTest(t, "conformance-files/thesubdb/justified.mp4", "edc1981d6459c6111fe36205b4aff6c2")
}
