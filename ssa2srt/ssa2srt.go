package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/filter"
	"github.com/martinlindhe/subber/srt"
	"github.com/martinlindhe/subber/ssa"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file       = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	keepAds    = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	filterName = kingpin.Flag("filter", "Filter (none, caps, html).").Default("none").String()
)

func main() {
	// support -h for --help
	kingpin.Version("0.1.0")
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
		fmt.Println(err)
		os.Exit(1)
	}

	captions := ssa.ParseSsa(data)
	if !*keepAds {
		captions = caption.CleanSubs(captions)
	}

	captions = filter.FilterSubs(captions, *filterName)

	srt.WriteSrt(captions, subFileName)

	fmt.Printf("Written %s\n", subFileName)
}
