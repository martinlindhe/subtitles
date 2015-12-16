package subber

import (
	"fmt"
	"io/ioutil"
	"time"
)

const tempFilePrefix = "moviehash-temp"

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func makeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
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
