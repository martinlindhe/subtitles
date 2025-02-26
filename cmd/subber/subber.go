package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/alecthomas/kingpin/v2"
	"github.com/martinlindhe/subtitles"
	log "github.com/sirupsen/logrus"
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
	filterName  = kingpin.Flag("filter", "Filter (none, caps, html, ocr, merge, flip, all).").Default("none").String()
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
		fmt.Printf("An error occurred processing '%s': %s\n", inFileName, err)
		os.Exit(1)
	}
}

func verboseMessage(args ...interface{}) {
	if !*quiet {
		fmt.Println(args...)
	}
}

func parseAndWriteSubFile(inFileName, outFileName string, filterName string, keepAds bool, sync int) error {
	data, err := os.ReadFile(inFileName)
	if err != nil {
		return err
	}

	out, err := cleanupSub(data, filterName, keepAds, sync, inFileName)
	if err != nil {
		return err
	}
	if string(data) == out {
		// only write to disk if content has changed
		return nil
	}
	return writeText(outFileName, *skipBackups, out)
}

// cleanupSub parses .srt or .ssa, performs cleanup and renders to a .srt, returning a string. caller is responsible for passing UTF8 string
func cleanupSub(data []byte, filterName string, keepAds bool, sync int, cleanerOutputPrefix string) (string, error) {
	subtitle, err := subtitles.Parse(data)
	if err != nil {
		return "", err
	}
	if !keepAds {
		subtitle.RemoveAds(cleanerOutputPrefix, true)
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
		var err error
		if !*dontTouch {
			err = parseAndWriteSubFile(inFileName, outFileName, *filterName, *keepAds, *sync)
		}
		return err
	}

	subFileName := inFileName[:len(inFileName)-len(ext)] + ".srt"

	if fileExists(subFileName) {
		fmt.Println("ERROR: Subs found locally in", subFileName, " (BUT DID NOT LOOK LIKE SUBS), skipping download")
		os.Exit(1)
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

	if *dontTouch {
		_, err = cleanupSub(data, *filterName, *keepAds, *sync, inFileName)
		if err != nil {
			log.Printf("ERROR: cleanupSub failed: %s", err)
		}

		// write untouched copy
		return writeText(subFileName, *skipBackups, string(data))
	}

	out, err := cleanupSub(data, *filterName, *keepAds, *sync, inFileName)
	if err != nil {
		return err
	}
	return writeText(subFileName, *skipBackups, out)
}

func writeText(outFileName string, skipBackups bool, text string) error {
	backupFileName := outFileName + ".org"
	if fileExists(outFileName) {
		err := os.Rename(outFileName, backupFileName)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}

	if skipBackups && fileExists(backupFileName) {
		err = os.Remove(backupFileName)
	}
	return err
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
