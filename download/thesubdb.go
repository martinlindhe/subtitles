package download

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func FromTheSubDb(fileName string) (string, error) {
	hash := createMovieHashFromMovieFile(fileName)

	actualText, err := downloadSubtitleByHash(hash)
	if err != nil {
		return actualText, err
	}

	return actualText, nil
}

// returns a md5-sum in hex-string representation
func createMovieHashFromMovieFile(fileName string) string {
	// block size which is required for the API call
	readSize := int64(64 * 1024)

	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	fi, err := f.Stat()
	check(err)

	if fi.Size() < readSize {
		fmt.Println("Input file is too small")
		return ""
	}

	// read first part
	b1 := make([]byte, readSize)
	_, err = f.Read(b1)
	check(err)

	// move the file pointer ahead, because we only need
	// the first and the last 64KB of the video file
	_, err = f.Seek(-readSize, 2)
	check(err)

	// read the last part
	b2 := make([]byte, readSize)
	_, err = f.Read(b2)
	check(err)

	combined := append(b1, b2...)

	return fmt.Sprintf("%x", md5.Sum(combined))
}

func downloadSubtitleByHash(hash string) (string, error) {

	client := &http.Client{}

	language := "en"

	query := "http://api.thesubdb.com/?action=download&hash=" + hash + "&language=" + language

	fmt.Printf("Fetching %s ...\n", query)

	req, err := http.NewRequest("GET", query, nil)
	check(err)

	req.Header.Add("User-Agent", "SubDB/1.0 (GoSubber/1.0; http://github.com/martinlindhe/go-subber)")
	resp, err := client.Do(req)
	check(err)

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
