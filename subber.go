package main

import (
	"fmt"
	"os"
	"path"

	"github.com/martinlindhe/go-subber/download"
	"github.com/martinlindhe/go-subber/srt"
	"gopkg.in/alecthomas/kingpin.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

var (
	file    = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	keepAds = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	inFileName := (*file).Name()

	if len(inFileName) < 1 {
		fmt.Printf("File name required\n")
		os.Exit(0)
	}

	ext := path.Ext(inFileName)
	if ext == ".srt" {
		srt.CleanupSrt(inFileName, true, !*keepAds)
		os.Exit(0)
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if exists(subFileName) {
		fmt.Println("Subs found locally, not downloading ...")
		srt.CleanupSrt(subFileName, true, !*keepAds)
		os.Exit(0)
	}

	fmt.Printf("Downloading subs for input file ...\n")

	text, err := download.FromTheSubDb(inFileName)
	check(err)

	f, err := os.Create(subFileName)
	check(err)

	f.WriteString(text)
	f.Close()

	srt.CleanupSrt(subFileName, true, !*keepAds)

}
