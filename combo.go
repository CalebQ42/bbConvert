package bbConvert

import (
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

// Convert BBCode and Markdown to HTML
func (c ComboConverter) HTMLConvert(combo string) string {
	//TODO: make this a bit more custom to prevent a couple, rare, collisions
	//_blank link target is the one I know of right now.
	combo = c.bb.HTMLConvert(combo)
	in := []rune(combo)
	var codeBlocks []string
	var match *regexp2.Match
	var err error
	// Pull back out any code blocks
	for {
		match, err = c.code.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		codeBlocks = append(codeBlocks, match.String())
		in = slices.Concat(in[:match.Index], []rune(codePlaceholder), in[match.Index+match.Length:])
	}
	out := c.md.HTMLConvert(string(in))
	for i := range codeBlocks {
		out = strings.Replace(out, codePlaceholder, codeBlocks[i], 1)
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
