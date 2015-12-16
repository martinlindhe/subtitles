package subber

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FindSub finds subtitle online, returns untouched data
func FindSub(videoFileName string, language string) ([]byte, error) {
	if !exists(videoFileName) {
		return nil, fmt.Errorf("%s not found", videoFileName)
	}

	if isDirectory(videoFileName) {
		return nil, fmt.Errorf("%s is not a file", videoFileName)
	}

	text, err := fromTheSubDb(videoFileName, language)
	if err != nil {
		return nil, err
	}

	return text, nil
}

// FromTheSubDb downloads a subtitle from thesubdb.com
func fromTheSubDb(videoFileName string, language string, optional ...string) ([]byte, error) {

	_apiHost := "api.thesubdb.com"
	if len(optional) > 0 {
		_apiHost = optional[0]
	}

	hash, err := createMovieHashFromMovieFile(videoFileName)
	if err != nil {
		return nil, err
	}

	actualText, err := downloadSubtitleByHash(hash, language, _apiHost)
	if err != nil {
		return nil, err
	}

	return actualText, nil
}

// returns a md5-sum in hex-string representation
func createMovieHashFromMovieFile(fileName string) (string, error) {

	if !exists(fileName) {
		return "", fmt.Errorf("File %s not found", fileName)
	}

	// block size which is required for the API call
	readSize := int64(64 * 1024)

	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	if fi.Size() < readSize {
		return "", fmt.Errorf("File is too small: %s", fileName)
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

func downloadSubtitleByHash(hash string, language string, apiHost string) ([]byte, error) {

	client := &http.Client{}

	query := "http://" + apiHost + "/?action=download&hash=" + hash + "&language=" + language

	log.Printf("Fetching %s ...\n", query)

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent",
		"SubDB/1.0 (GoSubber/1.0; http://github.com/martinlindhe/subber)")

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
