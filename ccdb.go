package subtitles

// an obscure (to me) text-based subtitle format, can have extension .cc or .txt
// found in The Way We Live Now (2001) BBC TV 1.txt
// STATUS: incomplete detection and support

import (
	"log"
	"strings"
	"time"
)

func looksLikeCCDBCapture(s string) bool {
	return strings.Contains(s, "[SUBTITLE]")
}

// NewFromCCDBCapture parses a ccdb capture text into []Caption, assumes s is a clean utf8 string
func NewFromCCDBCapture(s string) (res Subtitle, err error) {
	rows := strings.Split(s, "\n")
	seq := 1
	caption := Caption{Seq: seq}
	parseText := false
	for rowNum, row := range rows {
		if len(row) > 1 && row[0] == '[' {
			continue
		}
		if parseText {
			if row == "\r" || row == "" {
				parseText = false
			} else {
				row = strings.TrimSpace(row)
				if row != "" {
					caption.Text = append(caption.Text, row)
				}
			}
			if strings.Join(caption.Text, "") != "" {
				res.Captions = append(res.Captions, caption)
				seq++
				caption = Caption{Seq: seq}
			}
		} else if !parseText {
			if row == "" {
				if rowNum != len(rows)-1 {
					log.Println("NOTICE: ccdb seem to have reached end of valid stream at row", rowNum, "of", len(rows))
				}
				break
			}
			parts := strings.SplitN(row, ",", 2)
			if len(parts) == 2 {
				caption.Start, _ = parseCCDBTime(parts[0])
				caption.End, _ = parseCCDBTime(parts[1])
			} else {
				log.Println("TIME seq", seq, ", input row", (rowNum + 1), "error:", row)
			}
			parseText = true
		}
	}
	return
}

func parseCCDBTime(s string) (time.Time, error) {
	return parseTime(s)
}
