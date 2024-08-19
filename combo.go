package bbConvert

import (
	"strings"

	"github.com/dlclark/regexp2"
)

const CodeBlockRegEx = `<code>[\s\S]*?<\/code>`

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
		code: regexp2.MustCompile(CodeBlockRegEx, regexp2.Multiline),
	}
}

// Convert BBCode and Markdown to HTML
func (c ComboConverter) HTMLConvert(in string) string {
	in = c.bb.HTMLConvert(in)
	var codeBlocks []string
	var match *regexp2.Match
	var err error
	// Pull back out any code blocks
	for {
		match, err = c.code.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		codeBlocks = append(codeBlocks, match.String())
		in = in[:match.Index] + codePlaceholder + in[match.Index+match.Length:]
	}
	in = c.md.HTMLConvert(in)
	for i := range codeBlocks {
		in = strings.Replace(in, codePlaceholder, codeBlocks[i], 1)
	}
	return in
}

// Converts just BBCode to HTML
func (c ComboConverter) BBHTMLConvert(in string) string {
	return c.bb.HTMLConvert(in)
}

// Converts just Markdown to HTML
func (c ComboConverter) MarkdownHTMLConvert(in string) string {
	return c.md.HTMLConvert(in)
}
