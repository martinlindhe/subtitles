package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/martinlindhe/subber/cleaner"
	"github.com/martinlindhe/subber/download"
	"github.com/martinlindhe/subber/filter"
	"github.com/martinlindhe/subber/helpers"
	"github.com/martinlindhe/subber/parser"
	"github.com/martinlindhe/subber/srt"
	"github.com/martinlindhe/subber/txtformat"
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

const version = "0.1.2"

func main() {
	// support -h for --help
	kingpin.Version(version)
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
		cleanupSrt(inFileName, *filterName, *skipBackups, *keepAds)
		return nil
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if helpers.Exists(subFileName) {
		fmt.Printf("Subs found locally in %s, skipping download\n", subFileName)
		cleanupSrt(subFileName, *filterName, *skipBackups, *keepAds)
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

		captions := parser.Parse(data)

		if !*keepAds {
			captions = cleaner.RemoveAds(captions)
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

// CleanupSrt performs cleanup on fileName, overwriting the original file
func cleanupSrt(inFileName string, filterName string, skipBackup bool, keepAds bool) error {

	fmt.Fprintf(os.Stderr, "CleanupSrt %s\n", inFileName)

	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		return err
	}

	utf8 := txtformat.ConvertToUTF8(data)

	captions := srt.ParseSrt(utf8)
	if !keepAds {
		captions = cleaner.RemoveAds(captions)
	}

	captions = filter.FilterSubs(captions, filterName)

	out := srt.RenderSrt(captions)

	if string(data) == out {
		return nil
	}

	if !skipBackup {
		backupFileName := inFileName + ".org"
		os.Rename(inFileName, backupFileName)
		// fmt.Printf("Backed up to %s\n", backupFileName)
	}

	f, err := os.Create(inFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(out)
	if err != nil {
		return err
	}

	//fmt.Printf("Written %d captions to %s\n", len(captions), inFileName)
	return nil
}

func writeText(outFileName, text string) error {
	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}
