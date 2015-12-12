package download

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/martinlindhe/go-subber/common"
	"github.com/martinlindhe/go-subber/srt"
)

// FindSubs tries to find subtitles online
func FindSubs(videoFileName string, keepAds bool) ([]srt.Caption, error) {

	if !common.Exists(videoFileName) {
		return nil, fmt.Errorf("%s not found", videoFileName)
	}

	if common.IsDirectory(videoFileName) {
		return nil, fmt.Errorf("%s is not a file", videoFileName)
	}

	text, err := fromTheSubDb(videoFileName)
	if err != nil {
		return nil, err
	}

	captions := srt.ParseSrt(text)

	if !keepAds {
		captions = srt.CleanSubs(captions)
	}

	return captions, nil
}

// FromTheSubDb downloads a subtitle from thesubdb.com
func fromTheSubDb(videoFileName string, optional ...string) (string, error) {

	_apiHost := "api.thesubdb.com"
	if len(optional) > 0 {
		_apiHost = optional[0]
	}

	hash, err := createMovieHashFromMovieFile(videoFileName)
	if err != nil {
		return "", err
	}

	actualText, err := downloadSubtitleByHash(hash, _apiHost)
	if err != nil {
		return "", err
	}

	return actualText, nil
}

// returns a md5-sum in hex-string representation
func createMovieHashFromMovieFile(fileName string) (string, error) {

	// XXX make sure filename is a file, and not a dir
	if !common.Exists(fileName) {
		return "", fmt.Errorf("File %s not found", fileName)
	}

	// block size which is required for the API call
	readSize := int64(64 * 1024)

	f, err := os.Open(fileName)
	common.Check(err)
	defer f.Close()

	fi, err := f.Stat()
	common.Check(err)

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

func downloadSubtitleByHash(hash string, apiHost string) (string, error) {

	client := &http.Client{}

	language := "en"

	query := "http://" + apiHost + "/?action=download&hash=" + hash + "&language=" + language

	fmt.Printf("Fetching %s ...\n", query)

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent",
		"SubDB/1.0 (GoSubber/1.0; http://github.com/martinlindhe/go-subber)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 404 {
		return "", fmt.Errorf("Subtitle not found")
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Server error %s", resp.Status)
	}

	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Server reading request body: %v", err)
	}

	return string(slurp), nil
}
