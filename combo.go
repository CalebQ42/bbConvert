package bbConvert

import (
	"cmp"
	"slices"
	"strings"

	"github.com/dlclark/regexp2"
)

// A BBCode and Markdown converter all in one.
// Attemps to prevent some potential errors when doing them separately.
type ComboConverter struct {
	bb   BBConverter
	md   MarkdownConverter
	code *regexp2.Regexp
}

func NewComboConverter() ComboConverter {
	return ComboConverter{
		bb:   NewBBConverter(),
		md:   NewMarkdownConverter(),
		code: regexp2.MustCompile(`<code>[\s\S]*?<\/code>`, regexp2.Multiline),
	}
}

type codeMatch struct {
	index int
	code  string
}

// Convert BBCode and Markdown to HTML
func (c ComboConverter) HTMLConvert(combo string) string {
	in := []rune(combo)
	var codeBlocks []codeMatch
	var match *regexp2.Match
	var err error
	var ind, dif int
	for {
		match, err = c.bb.codeConv.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		if strings.Contains(match.GroupByNumber(1).String(), "\n") {
			in = slices.Concat(in[:match.Index], []rune("</p><pre>"+codePlaceholder+"</pre><p>"), in[match.Index+match.Length:])
		} else {
			in = slices.Concat(in[:match.Index], []rune(codePlaceholder), in[match.Index+match.Length:])
		}
		n := codeMatch{match.Index, strings.TrimPrefix(match.GroupByNumber(1).String(), "\n")}
		dif = len(n.code) - len(codePlaceholder)
		ind, _ = slices.BinarySearchFunc(codeBlocks, n, func(a, b codeMatch) int { return cmp.Compare(a.index, b.index) })
		for i := ind; i < len(codeBlocks); i++ {
			codeBlocks[i].index -= dif
		}
		codeBlocks = slices.Insert(codeBlocks, ind, n)
	}
	for {
		match, err = c.md.largeCodeConv.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		in = slices.Concat(in[:match.Index], []rune("</p><pre>"+codePlaceholder+"</pre><p>"), in[match.Index+match.Length:])
		n := codeMatch{match.Index, strings.TrimPrefix(match.GroupByNumber(1).String(), "\n")}
		dif = len(n.code) - len(codePlaceholder)
		ind, _ = slices.BinarySearchFunc(codeBlocks, n, func(a, b codeMatch) int { return cmp.Compare(a.index, b.index) })
		for i := ind; i < len(codeBlocks); i++ {
			codeBlocks[i].index -= dif
		}
		codeBlocks = slices.Insert(codeBlocks, ind, n)
	}
	for {
		match, err = c.md.inlineCodeConv.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		in = slices.Concat(in[:match.Index], []rune(codePlaceholder), in[match.Index+match.Length:])
		n := codeMatch{match.Index, strings.TrimPrefix(match.GroupByNumber(1).String(), "\n")}
		dif = len(n.code) - len(codePlaceholder)
		ind, _ = slices.BinarySearchFunc(codeBlocks, n, func(a, b codeMatch) int { return cmp.Compare(a.index, b.index) })
		for i := ind; i < len(codeBlocks); i++ {
			codeBlocks[i].index -= dif
		}
		codeBlocks = slices.Insert(codeBlocks, ind, n)
	}
	out := c.bb.bbActualConv([]rune(c.md.mkActualConv(in, true)), true)
	for i := range codeBlocks {
		out = strings.Replace(out, codeInner, codeBlocks[i].code, 1)
	}
	return out
}

// Converts just BBCode to HTML
func (c ComboConverter) BBHTMLConvert(in string) string {
	return c.bb.HTMLConvert(in)
}

// Converts just Markdown to HTML
func (c ComboConverter) MarkdownHTMLConvert(in string) string {
	return c.md.HTMLConvert(in)
}
