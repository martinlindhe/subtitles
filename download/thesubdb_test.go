package download

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createZeroedTempFile(byteSize int) string {
	emptyData := make([]byte, byteSize)

	f, err := ioutil.TempFile("/tmp", "moviehash-temp")
	check(err)

	defer f.Close()

	fileName := f.Name()

	f.Write(emptyData)

	return fileName
}

func TestCreateMovieHashFromMovieFile(t *testing.T) {

	fileName := createZeroedTempFile(1024 * 1024 * 2)

	assert.Equal(t, "0dfbe8aa4c20b52e1b8bf3cb6cbdf193", createMovieHashFromMovieFile(fileName))

	os.Remove(fileName)
}

func TestDownloadFromTheSubDb(t *testing.T) {

	fileName := createZeroedTempFile(1024 * 1024 * 2)

	assert.Equal(t, "", FromTheSubDb(fileName))

	os.Remove(fileName)
}
