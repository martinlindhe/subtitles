package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func MakeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
}

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDirectory reports wether the named path is a directory
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func CreateTempFile(byteSize int) string {
	data := make([]byte, byteSize)

	cnt := uint8(0)
	for i := 0; i < byteSize; i++ {
		data[i] = cnt
		cnt++
	}

	f, err := ioutil.TempFile("/tmp", "moviehash-temp")
	Check(err)

	defer f.Close()

	fileName := f.Name()

	f.Write(data)

	return fileName
}

func CreateZeroedTempFile(byteSize int) string {
	data := make([]byte, byteSize)

	f, err := ioutil.TempFile("/tmp", "moviehash-temp")
	Check(err)

	defer f.Close()

	fileName := f.Name()

	f.Write(data)

	return fileName
}
