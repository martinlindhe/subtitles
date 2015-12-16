package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/martinlindhe/subber"
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

const version = "0.2.0"

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

		data, err := subber.ReadFileAsUTF8(inFileName)
		if err != nil {
			return err
		}

		out, err := subber.CleanupSub(data, *filterName, *keepAds)
		if err != nil {
			return err
		}
		writeText(inFileName, *skipBackups, out)
		return nil
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if fileExists(subFileName) {
		fmt.Printf("Subs found locally in %s, skipping download\n", subFileName)

		data, err := subber.ReadFileAsUTF8(subFileName)
		if err != nil {
			return err
		}

		out, err := subber.CleanupSub(data, *filterName, *keepAds)
		if err != nil {
			return err
		}
		writeText(inFileName, *skipBackups, out)
		return nil
	}

	fmt.Printf("Downloading subs for %s ...\n", inFileName)

	data, err := subber.FindSub(inFileName, *language)
	if err != nil {
		return err
	}

	out := ""

	if *dontTouch {
		out = string(data)
		// write untouched copy
		err = writeText(subFileName+".org", *skipBackups, out)
		if err != nil {
			return err
		}
	} else {
		out, _ = subber.CleanupSub(subFileName+".org", *filterName, *keepAds)
	}

	err = writeText(subFileName, *skipBackups, out)
	if err != nil {
		return err
	}

	return nil
}

func writeText(outFileName string, skipBackups bool, text string) error {

	if !skipBackups && fileExists(outFileName) {
		backupFileName := outFileName + ".org"
		os.Rename(outFileName, backupFileName)
		// fmt.Printf("Backed up to %s\n", backupFileName)
	}

	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
