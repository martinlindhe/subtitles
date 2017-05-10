package subtitles

import (
	"fmt"
	"io/ioutil"
)

const tempFilePrefix = "moviehash-temp"

func check(e error) {

	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func createTempFile(byteSize int) string {

	data := make([]byte, byteSize)

	cnt := uint8(0)
	for i := 0; i < byteSize; i++ {
		data[i] = cnt
		cnt++
	}

	f, err := ioutil.TempFile("/tmp", tempFilePrefix)
	check(err)
	defer f.Close()

	fileName := f.Name()
	f.Write(data)

	return fileName
}

func createZeroedTempFile(byteSize int) string {

	data := make([]byte, byteSize)

	f, err := ioutil.TempFile("/tmp", tempFilePrefix)
	check(err)
	defer f.Close()

	fileName := f.Name()
	f.Write(data)

	return fileName
}
