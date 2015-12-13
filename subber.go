package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/martinlindhe/go-subber/download"
	"github.com/martinlindhe/go-subber/helpers"
	"github.com/martinlindhe/go-subber/srt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file        = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	dontTouch   = kingpin.Flag("dont-touch", "Do not try to process .srt (write directly to disk).").Bool()
	keepAds     = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	skipBackups = kingpin.Flag("skip-backups", "Do not make backup (.srt.org) of original .srt").Bool()
	language    = kingpin.Flag("language", "Language.").Default("en").String()
	filter      = kingpin.Flag("filter", "Filter (none, caps, html).").Default("none").String()
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

	err := action(inFileName)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

func action(inFileName string) error {

	ext := path.Ext(inFileName)
	if ext == ".srt" {
		srt.CleanupSrt(inFileName, *filter, *skipBackups, *keepAds)
		return nil
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if helpers.Exists(subFileName) {
		fmt.Printf("Subs found locally in %s, skipping download\n", subFileName)
		srt.CleanupSrt(subFileName, *filter, *skipBackups, *keepAds)
		return nil
	}

	fmt.Printf("Downloading subs for %s ...\n", inFileName)

	if *dontTouch {
		// download and write untouched
		text, err := download.FindSubText(inFileName, *language)
		if err != nil {
			return err
		}

		err = writeText(subFileName, text)
		if err != nil {
			return err
		}
	} else {

		captions, err := download.FindSub(inFileName, *language, *keepAds)
		if err != nil {
			return err
		}

		err = srt.WriteSrt(captions, subFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeText(outFileName, text string) error {
	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}
