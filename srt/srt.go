package srt

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Eol is the end of line characters to use when writing .srt data
const Eol = "\n"

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// Caption represents one subtitle block
type Caption struct {
	seq   int
	start time.Time
	end   time.Time
	text  []string
}

func renderSrtTime(t time.Time) string {
	res := t.Format("15:04:05.000")
	return strings.Replace(res, ".", ",", 1)
}

func (cap Caption) srtTime() string {
	return renderSrtTime(cap.start) + " --> " + renderSrtTime(cap.end)
}

// ParseSrt parses a .srt text into []Caption
func ParseSrt(s string) []Caption {

	var res []Caption

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

		var o Caption
		o.seq = val
		i++

		matches := r1.FindStringSubmatch(lines[i])

		o.start, err = parseTime(matches[1])
		if err != nil {
			fmt.Printf("Parse error at line %d: %v\n", i, err)
			continue
		}

		o.end, err = parseTime(matches[2])
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
			o.text = append(o.text, line)
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
func WriteSrt(subs []Caption, outFileName string) {

	text := RenderSrt(subs)

	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		fmt.Printf("Error writing to %s, %v", outFileName, err)
	}
}

// RenderSrt produces a text representation of the subtitles
func RenderSrt(subs []Caption) string {

	res := ""

	for _, sub := range subs {
		res += renderCaptionAsSrt(sub)
	}

	return res
}

func renderCaptionAsSrt(sub Caption) string {

	res := fmt.Sprintf("%d", sub.seq) + Eol +
		sub.srtTime() + Eol

	for _, line := range sub.text {
		res += line + Eol
	}

	return res + Eol
}

// CleanupSrt performs cleanup on fileName, overwriting the original file
func CleanupSrt(inFileName string, makeBackup bool) {

	fmt.Printf("Cleaning sub %s ...\n", inFileName)

	data, err := ioutil.ReadFile(inFileName)
	check(err)

	s := string(data)

	subs := ParseSrt(s)

	cleaned := CleanSubs(subs)

	out := RenderSrt(cleaned)

	if s == out {
		fmt.Printf("No changes performed\n")
		return
	}

	if makeBackup {
		backupFileName := inFileName + ".org"
		os.Rename(inFileName, backupFileName)
		fmt.Printf("Backed up to %s\n", backupFileName)
	}

	f, err := os.Create(inFileName) // xxx can we create if exists? when makebackup=false ?
	check(err)
	defer f.Close()

	_, err = f.WriteString(out)
	check(err)

	fmt.Printf("Written to %s\n", inFileName)
}
