package srt

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Caption represents one subtitle block
type Caption struct {
	seq   int
	start time.Time
	end   time.Time
	text  []string
}

// ParseSrt parses a .srt text representation into []Caption
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
			fmt.Printf("Parse error at line %d: ", i)
			fmt.Println(err)
			continue
		}

		var o Caption
		o.seq = val
		i++

		matches := r1.FindStringSubmatch(lines[i])

		o.start = parseTime(matches[1])
		o.end = parseTime(matches[2])
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

func parseTime(in string) time.Time {

	// . to ,
	in = strings.Replace(in, ",", ".", 1)

	if !strings.ContainsAny(in, ".") {
		in += ".000"
	}

	const form = "15:04:05.000"
	t, err := time.Parse(form, in)
	if err != nil {
		fmt.Println(err)
	}

	return t
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
	res := fmt.Sprintf("%d\n", sub.seq) + sub.srtTime() + "\n"

	for _, line := range sub.text {
		res += line + "\n"
	}

	res += "\n"

	return res
}

func renderSrtTime(t time.Time) string {
	const form = "15:04:05.000"
	res := t.Format(form)

	return strings.Replace(res, ".", ",", 1)
}

func (cap Caption) srtTime() string {
	return renderSrtTime(cap.start) + " --> " + renderSrtTime(cap.end)
}
