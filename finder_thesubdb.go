package subtitles

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TheSubDb downloads a subtitle from thesubdb.com
func (f SubFinder) TheSubDb(args ...string) ([]byte, error) {

	apiHost := ""
	if len(args) > 0 {
		apiHost = args[0]
	} else {
		apiHost = "api.thesubdb.com"
	}

	hash, err := SubDbHashFromFile(f.VideoFile)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	query := "http://" + apiHost +
		"/?action=download" +
		"&hash=" + hash +
		"&language=" + f.Language

	if !f.Quiet {
		fmt.Println("Fetching", query, "...")
	}

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent",
		"SubDB/1.0 (GoSubber/1.0; https://github.com/martinlindhe/subtitles)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Subtitle not found")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Server error %s", resp.Status)
	}

	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Server reading request body: %v", err)
	}

	return slurp, nil
}

// SubDbHashFromFile returns a checksum in hex-string representation
// conforming to http://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes
func SubDbHashFromFile(f *os.File) (string, error) {

	// rewind
	f.Seek(0, 0)

	// block size which is required for the API call
	readSize := int64(64 * 1024)

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	if fi.Size() < readSize {
		return "", fmt.Errorf("Stream is too small: %d", fi.Size())
	}

	// read first part
	b1 := make([]byte, readSize)
	_, err = f.Read(b1)
	if err != nil {
		return "", err
	}

	// move the file pointer ahead, because we only need
	// the first and the last 64KB of the video file
	_, err = f.Seek(-readSize, 2)
	if err != nil {
		return "", err
	}

	// read the last part
	b2 := make([]byte, readSize)
	_, err = f.Read(b2)
	if err != nil {
		return "", err
	}

	combined := append(b1, b2...)

	return fmt.Sprintf("%x", md5.Sum(combined)), nil
}
