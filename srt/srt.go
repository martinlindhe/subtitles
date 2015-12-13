package srt

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/filter"
	"github.com/martinlindhe/subber/txtformat"
)

// Eol is the end of line characters to use when writing .srt data
const Eol = "\n"

// ParseSrt parses a .srt text into []Caption
func ParseSrt(b []byte) []caption.Caption {

	s := txtformat.ConvertToUTF8(b)

	var res []caption.Caption

	r1 := regexp.MustCompile("([0-9:.,]*) --> ([0-9:.,]*)")

	lines := strings.Split(s, "\n")

	outSeq := 1

	for i := 0; i < len(lines); i++ {

		seq := strings.Trim(lines[i], "\r ")
		if seq == "" {
			break
		}

		_, err := strconv.Atoi(seq)
		if err != nil {
			fmt.Printf("Parse error 1 at line %d: %v\n", i, err)
			continue
		}

		var o caption.Caption
		o.Seq = outSeq
		i++
		if i >= len(lines) {
			break
		}

		matches := r1.FindStringSubmatch(lines[i])
		if len(matches) < 3 {
			fmt.Printf("Parser error 2 at line %d (idx out of range)\n", i)
			continue
		}

		o.Start, err = parseTime(matches[1])
		if err != nil {
			fmt.Printf("Parse error 3 at line %d: %v\n", i, err)
			continue
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			fmt.Printf("Parse error 4 at line %d: %v\n", i, err)
			continue
		}

		i++

		textLine := 1
		for {
			line := strings.Trim(lines[i], "\r ")
			if line == "" && textLine > 1 {
				break
			}
			if line != "" {
				o.Text = append(o.Text, line)
			}
			i++

			if i >= len(lines) {
				break
			}
			textLine++
		}

		if len(o.Text) > 0 {
			res = append(res, o)
			outSeq++
		}
	}

	return res
}

func makeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
}

func parseTime(in string) (time.Time, error) {

	// . to ,
	in = strings.Replace(in, ",", ".", 1)

	if !strings.ContainsAny(in, ".") {
		in += ".000"
	}

	r1 := regexp.MustCompile("([0-9]+):([0-9]+):([0-9]+)[.]([0-9]+)")

	matches := r1.FindStringSubmatch(in)

	if len(matches) < 5 {
		return time.Now(), fmt.Errorf("Regexp didnt match: %s", in)
	}

	h, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Now(), err
	}

	m, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Now(), err
	}

	s, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Now(), err
	}

	ms, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Now(), err
	}

	return makeTime(h, m, s, ms), nil
}

// WriteSrt prints a srt render to outFileName
func WriteSrt(subs []caption.Caption, outFileName string) error {

	text := RenderSrt(subs)

	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
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

	fmt.Fprintf(os.Stderr, "CleanupSrt %s\n", inFileName)

	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		return err
	}

	captions := ParseSrt(data)
	if !keepAds {
		captions = caption.CleanSubs(captions)
	}

	captions = filter.FilterSubs(captions, filterName)

	out := RenderSrt(captions)

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
