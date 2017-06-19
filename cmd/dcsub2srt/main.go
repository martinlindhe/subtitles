// XXX should just invoke "subber vid.xml -o vid.srt"

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/subtitles"
)

var (
	name = kingpin.Arg("name", "DCSubtitle file.").Required().String()
)

func main() {
	kingpin.Parse()

	data, err := ioutil.ReadFile(*name)
	if err != nil {
		log.Fatal(err)
	}
	sub, err := subtitles.NewFromDCSub(string(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(sub.AsSRT())
}
