package srt

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/martinlindhe/go-subber/caption"
	"github.com/martinlindhe/go-subber/filter"
)

// Eol is the end of line characters to use when writing .srt data
const Eol = "\n"

// ParseSrt parses a .srt text into []Caption
func ParseSrt(s string) []caption.Caption {

	var res []caption.Caption

	r1 := regexp.MustCompile("([0-9:,]*) --> ([0-9:,]*)")

	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i++ {

		seq := strings.Trim(lines[i], "\ufeff\r ")
		if seq == "" {
			break
		}

		val, err := strconv.Atoi(seq)
		if err != nil {
			fmt.Printf("Parse error at line %d: %v\n", i, err)
			continue
		}

		var o caption.Caption
		o.Seq = val
		i++

		matches := r1.FindStringSubmatch(lines[i])

		o.Start, err = parseTime(matches[1])
		if err != nil {
			fmt.Printf("Parse error at line %d: %v\n", i, err)
			continue
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			fmt.Printf("Parse error at line %d: %v\n", i, err)
			continue
		}

		i++

		for {
			line := strings.Trim(lines[i], "\r ")
			if i >= len(lines) || line == "" {
				break
			}
			o.Text = append(o.Text, line)
			i++
		}

		res = append(res, o)
	}

	return res
}

func parseTime(in string) (time.Time, error) {

	// . to ,
	in = strings.Replace(in, ",", ".", 1)

	if !strings.ContainsAny(in, ".") {
		in += ".000"
	}

	const form = "15:04:05.000"
	t, err := time.Parse(form, in)
	if err != nil {
		return t, fmt.Errorf("Parse error: %v", err)
	}

	return t, nil
}

// WriteSrt prints a srt render to outFileName
func WriteSrt(subs []caption.Caption, outFileName string) {

	text := RenderSrt(subs)

	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		fmt.Printf("Error writing to %s, %v", outFileName, err)
	}
}

// RenderSrt produces a text representation of the subtitles
func RenderSrt(subs []caption.Caption) string {

	res := ""

	for _, sub := range subs {
		res += renderCaptionAsSrt(sub)
	}

	return res
}

func renderCaptionAsSrt(caption caption.Caption) string {

	res := fmt.Sprintf("%d", caption.Seq) + Eol +
		caption.SrtTime() + Eol

	for _, line := range caption.Text {
		res += line + Eol
	}

	return res + Eol
}

// CleanupSrt performs cleanup on fileName, overwriting the original file
func CleanupSrt(inFileName string, filterName string, skipBackup bool, keepAds bool) error {

	fmt.Printf("Cleaning sub %s ...\n", inFileName)

	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		return err
	}

	s := string(data)

	captions := ParseSrt(s)
	if !keepAds {
		captions = caption.CleanSubs(captions)
	}

	captions = filter.FilterSubs(captions, filterName)

	out := RenderSrt(captions)

	if s == out {
		fmt.Printf("No changes performed\n")
		return nil
	}

	if !skipBackup {
		backupFileName := inFileName + ".org"
		os.Rename(inFileName, backupFileName)
		fmt.Printf("Backed up to %s\n", backupFileName)
	}

	f, err := os.Create(inFileName) // xxx can we create if exists? when makebackup=false ?
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(out)
	if err != nil {
		return err
	}

	fmt.Printf("Written %d captions to %s\n", len(captions), inFileName)
	return nil
}
