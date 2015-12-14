// +build subber

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/download"
	"github.com/martinlindhe/subber/helpers"
	"github.com/martinlindhe/subber/srt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file        = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	dontTouch   = kingpin.Flag("dont-touch", "Do not try to process .srt (write directly to disk).").Bool()
	keepAds     = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	skipBackups = kingpin.Flag("skip-backups", "Do not make backup (.srt.org) of original .srt").Bool()
	language    = kingpin.Flag("language", "Language.").Default("en").String()
	filterName  = kingpin.Flag("filter", "Filter (none, caps, html).").Default("none").String()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	inFileName := (*file).Name()

	// skip "hidden" .dotfiles
	baseName := filepath.Base(inFileName)
	if baseName[0] == '.' {
		// fmt.Printf("Skipping hidden %s\n", inFileName)
		os.Exit(1)
	}

	err := action(inFileName)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

func action(inFileName string) error {

	ext := path.Ext(inFileName)
	if ext == ".srt" {
		srt.CleanupSrt(inFileName, *filterName, *skipBackups, *keepAds)
		return nil
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if helpers.Exists(subFileName) {
		fmt.Printf("Subs found locally in %s, skipping download\n", subFileName)
		srt.CleanupSrt(subFileName, *filter, *skipBackups, *keepAds)
		return nil
	}

	fmt.Printf("Downloading subs for %s ...\n", inFileName)

	data, err := download.FindSubText(inFileName, *language)
	if err != nil {
		return err
	}

	if !*dontTouch {
		// write untouched copy
		err = writeText(subFileName+".org", string(data))
		if err != nil {
			return err
		}

		// clean and render to str
		captions := srt.ParseSrt(data)

		if !*keepAds {
			captions = caption.CleanSubs(captions)
		}

		text := srt.RenderSrt(captions)
		data = []byte(text)
	}

	err = writeText(subFileName, string(data))
	if err != nil {
		return err
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
