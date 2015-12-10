package main

import (
	"fmt"
	"os"
	"path"

	"github.com/martinlindhe/go-subber/download"
	"github.com/martinlindhe/go-subber/srt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file.srt>\n", os.Args[0])
		os.Exit(0)
	}

	inFileName := os.Args[1]

	ext := path.Ext(inFileName)
	if ext == ".srt" {

		srt.CleanupSrt(inFileName, true)
		os.Exit(0)
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if fileExists(subFileName) {
		fmt.Println("Subs found locally, not downloading ...")
		srt.CleanupSrt(subFileName, true)
		os.Exit(0)
	}

	fmt.Printf("Downloading subs for input file ...\n")

	text, err := download.FromTheSubDb(inFileName)
	check(err)

	f, err := os.Create(subFileName)
	check(err)

	f.WriteString(text)
	f.Close()

	srt.CleanupSrt(subFileName, true)
}
