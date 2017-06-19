// XXX should just invoke "subber vid.ssa -o vid.srt"

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/martinlindhe/subtitles"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file       = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	keepAds    = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	filterName = kingpin.Flag("filter", "Filter (none, caps, html).").Default("none").String()
)

const version = "0.1.1"

func main() {
	// support -h for --help
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	inFileName := (*file).Name()
	ext := path.Ext(inFileName)
	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	// skip "hidden" .dotfiles
	baseName := filepath.Base(inFileName)
	if baseName[0] == '.' {
		// fmt.Printf("Skipping hidden %s\n", inFileName)
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		log.Fatal(err)
	}

	out, err := cleanupSub(data, *filterName, *keepAds, 0)
	if err != nil {
		log.Fatal(err)
	}

	writeText(subFileName, false, out)

	fmt.Printf("Written %s\n", subFileName)
}

// cleanupSub parses .srt or .ssa, performs cleanup and renders to a .srt, returning a string. caller is responsible for passing UTF8 string
func cleanupSub(data []byte, filterName string, keepAds bool, sync int) (string, error) {
	subtitle, err := subtitles.Parse(data)
	if err != nil {
		return "", err
	}
	if !keepAds {
		subtitle.RemoveAds()
	}
	if sync != 0 {
		subtitle.ResyncSubs(sync)
	}
	subtitle.FilterCaptions(filterName)
	out := subtitle.AsSRT()
	return out, nil
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
