package download

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile(byteSize int) string {
	data := make([]byte, byteSize)

	cnt := uint8(0)
	for i := 0; i < byteSize; i++ {
		data[i] = cnt
		cnt++
	}

	f, err := ioutil.TempFile("/tmp", "moviehash-temp")
	check(err)

	defer f.Close()

	fileName := f.Name()

	f.Write(data)

	return fileName
}

func createZeroedTempFile(byteSize int) string {
	data := make([]byte, byteSize)

	f, err := ioutil.TempFile("/tmp", "moviehash-temp")
	check(err)

	defer f.Close()

	fileName := f.Name()

	f.Write(data)

	return fileName
}

func TestCreateMovieHashFromMovieFile(t *testing.T) {

	fileName := createTempFile(1024 * 1024 * 2)

	assert.Equal(t, "38a503307786991a982f8ded498b90e0", createMovieHashFromMovieFile(fileName))

	os.Remove(fileName)
}

func TestDownloadFromTheSubDb(t *testing.T) {
	fileName := createZeroedTempFile(1024 * 1024 * 2)

	text, err := FromTheSubDb(fileName, "sandbox.thesubdb.com")
	assert.Equal(t, nil, err)
	assert.True(t, len(text) > 1000)

	os.Remove(fileName)
}
