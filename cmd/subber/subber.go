package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/martinlindhe/subtitles"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file        = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	outFile     = kingpin.Flag("output-file", "A .srt (to clean) or video file (to fetch subs).").Short('o').String()
	verbose     = kingpin.Flag("verbose", "Verbose mode (more output).").Short('v').Bool()
	quiet       = kingpin.Flag("quiet", "Quiet mode (less output).").Short('q').Bool()
	dontTouch   = kingpin.Flag("dont-touch", "Do not try to process .srt (write directly to disk).").Bool()
	keepAds     = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	skipBackups = kingpin.Flag("skip-backups", "Do not make backup (.srt.org) of original .srt").Bool()
	language    = kingpin.Flag("language", "Language.").Default("en").String()
	filterName  = kingpin.Flag("filter", "Filter (none, caps, html, ocr, flip, all).").Default("none").String()
	sync        = kingpin.Flag("sync", "Synchronize captions (milliseconds).").Int()
)

const version = "0.3.0"

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(version)
	kingpin.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	inFileName := (*file).Name()
	outFileName := *outFile

	// skip "hidden" .dotfiles
	baseName := filepath.Base(inFileName)
	if baseName[0] == '.' {
		// fmt.Printf("Skipping hidden %s\n", inFileName)
		os.Exit(1)
	}

	err := action(inFileName, outFileName)
	if err != nil {
		fmt.Println("An error occured:", err)
	}
}

func verboseMessage(args ...interface{}) {
	if !*quiet {
		fmt.Println(args...)
	}
}

func parseAndWriteSubFile(inFileName, outFileName string, filterName string, keepAds bool, sync int) error {
	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		return err
	}
	out, err := cleanupSub(data, filterName, keepAds, sync)
	if err != nil {
		return err
	}
	return writeText(outFileName, *skipBackups, out)
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

func action(inFileName, outFileName string) error {
	if outFileName == "" {
		outFileName = inFileName
	}

	ext := path.Ext(inFileName)
	if subtitles.LooksLikeTextSubtitle(inFileName) {
		if !*dontTouch {
			parseAndWriteSubFile(inFileName, outFileName, *filterName, *keepAds, *sync)
		}
		return nil
	}

	subFileName := inFileName[:len(inFileName)-len(ext)] + ".srt"

	if fileExists(subFileName) {
		verboseMessage("Subs found locally in", subFileName, ", skipping download")
		if !*dontTouch {
			parseAndWriteSubFile(subFileName, outFileName, *filterName, *keepAds, *sync)
		}
		return nil
	}

	verboseMessage("Downloading subs for", inFileName, "...")

	f, err := os.Open(inFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	finder := subtitles.NewSubFinder(f, inFileName, *language)
	finder.Quiet = *quiet

	data, err := finder.TheSubDb("")
	if err != nil {
		return err
	}

	out := ""

	if *dontTouch {
		// write untouched copy
		err = writeText(subFileName, *skipBackups, string(data))
		if err != nil {
			return err
		}
		return nil
	}
	out, _ = cleanupSub(data, *filterName, *keepAds, *sync)
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
	return ioutil.WriteFile(outFileName, []byte(text), 0644)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
