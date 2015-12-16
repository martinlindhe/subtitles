package subber

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Eol is the end of line characters to use when writing .srt data
const eol = "\n"

func looksLikeSrt(s string) bool {
	if strings.HasPrefix(s, "1\n") || strings.HasPrefix(s, "1\r\n") {
		return true
	}
	return false
}

// ParseSrt parses a .srt text into []Caption, assumes s is a clean utf8 string
func parseSrt(s string) []caption {

	var res []caption

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
			fmt.Printf("[srt] Parse error 1 at line %d: %v\n", i, err)
			continue
		}

		var o caption
		o.Seq = outSeq

		i++
		if i >= len(lines) {
			break
		}

		matches := r1.FindStringSubmatch(lines[i])
		if len(matches) < 3 {
			fmt.Printf("[srt] Parser error 2 at line %d (idx out of range)\n", i)
			continue
		}

		o.Start, err = parseTime(matches[1])
		if err != nil {
			fmt.Printf("[srt] Parse error 3 at line %d: %v\n", i, err)
			continue
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			fmt.Printf("[srt] Parse error 4 at line %d: %v\n", i, err)
			continue
		}

		i++
		if i >= len(lines) {
			break
		}

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

func parseTime(in string) (time.Time, error) {

	// . to ,
	in = strings.Replace(in, ",", ".", 1)

	if !strings.ContainsAny(in, ".") {
		in += ".000"
	}

	r1 := regexp.MustCompile("([0-9]+):([0-9]+):([0-9]+)[.]([0-9]+)")

	matches := r1.FindStringSubmatch(in)

	if len(matches) < 5 {
		return time.Now(), fmt.Errorf("[srt] Regexp didnt match: %s", in)
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

// writeSrt prints a srt render to outFileName
func writeSrt(subs []caption, outFileName string) error {

	text := renderSrt(subs)

	err := ioutil.WriteFile(outFileName, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}

// renderSrt produces a text representation of the subtitles
func renderSrt(subs []caption) string {

	res := ""

	for _, sub := range subs {
		res += renderCaptionAsSrt(sub)
	}

	return res
}

func renderCaptionAsSrt(cap caption) string {

	res := fmt.Sprintf("%d", cap.Seq) + eol +
		cap.srtTime() + eol

	for _, line := range cap.Text {
		res += line + eol
	}

	return res + eol
}

func renderSrtTime(t time.Time) string {
	res := t.Format("15:04:05.000")
	return strings.Replace(res, ".", ",", 1)
}

func (cap caption) srtTime() string {
	return renderSrtTime(cap.Start) + " --> " + renderSrtTime(cap.End)
}
