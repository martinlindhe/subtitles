package subtitles

import (
	"bytes"
	"html"
	"io"
	"strconv"
	"strings"

	parser "golang.org/x/net/html"
)

type dcsubParser struct {
	cap  *Caption
	caps []*Caption
}

func looksLikeDCSub(s string) bool {
	return strings.Contains(strings.ToLower(s), "<dcsubtitle")
}

// NewFromDCSub parses a dcsubtitle xml sub, assumes s is a clean utf8 string
func NewFromDCSub(s string) (res Subtitle, err error) {
	r := strings.NewReader(s)
	var doc *parser.Node
	doc, err = parser.Parse(r)
	if err != nil {
		return
	}
	parser := dcsubParser{}
	sub, _ := parser.traverse(doc)
	return sub, nil
}

func (parser *dcsubParser) AsSubtitle() (res Subtitle) {
	for _, cap := range parser.caps {
		res.Captions = append(res.Captions, *cap)
	}
	return
}

func isSubtitleElement(n *parser.Node) bool {
	return n.Type == parser.ElementNode && n.Data == "subtitle"
}
func isTextElement(n *parser.Node) bool {
	return n.Type == parser.ElementNode && n.Data == "text"
}

func (parser *dcsubParser) traverse(n *parser.Node) (Subtitle, bool) {
	if isSubtitleElement(n) {
		parser.cap = &Caption{}
		parser.caps = append(parser.caps, parser.cap)
		if p, ok := getAttribute(n, "spotnumber"); ok {
			parser.cap.Seq, _ = strconv.Atoi(p)
		}
		if p, ok := getAttribute(n, "timein"); ok {
			parser.cap.Start, _ = parseTime(p)
		}
		if p, ok := getAttribute(n, "timeout"); ok {
			parser.cap.End, _ = parseTime(p)
		}
	}

	if isTextElement(n) {
		txt := renderNode(n.FirstChild)
		txt = html.UnescapeString(txt)
		parser.cap.Text = append(parser.cap.Text, txt)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := parser.traverse(c)
		if ok {
			return result, ok
		}
	}

	return parser.AsSubtitle(), false
}

func renderNode(n *parser.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	parser.Render(w, n)
	return buf.String()
}

func getAttribute(n *parser.Node, key string) (res string, ok bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			res = attr.Val
			ok = true
		}
	}
	return
}
