package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinlindhe/go-subber/srt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file.srt>\n", os.Args[0])
		os.Exit(0)
	}

	inFileName := os.Args[1]

	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := string(data)

	subs := srt.ParseSrt(s)

	cleaned := srt.CleanSubs(subs)

	out := srt.RenderSrt(cleaned)

	if s == out {
		fmt.Printf("No changes performed\n")
		os.Exit(0)
	}

	orgFileName := inFileName + ".org"
	os.Rename(inFileName, orgFileName)

	f, err := os.Create(inFileName)
	check(err)
	defer f.Close()

	_, err = f.WriteString(out)
	check(err)

	fmt.Print("Written to %s\n", inFileName)
}
