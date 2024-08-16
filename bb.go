package bbConvert

import (
	"strings"

	"github.com/dlclark/regexp2"
)

// The Magic RegEx string that matches bbCode tags.
const (
	BBMatchRegEx    = `\[(b|bold|i|italics|u|underline|s|strike|font|size|color|smallcaps|url|link|youtube|img|image|title|t[1-6]|align|float|ul|bullet|ol|number)(.*?)\]([\s\S]*?)\[\/\1\]`
	BBCustomRegEx   = `\[(\w+\b)(.*?)\]([\s\S]*?)\[\/\1\]`
	BBCodeRegEx     = `\[code\]([\s\S]*?)\[\/code\]`
	codePlaceholder = `%CODEBLOCK%`
)

type BBTag struct {
	Tag string
	// Leading value (if it exists). Eg. [img=500x200]
	Leading string
	// If the parameter doesn't have an associated value then it'll be set to an empty string
	// Eg. [float right] = {"right": ""}
	// Eg. [font color=red] = {"color": "red"}
	Values map[string]string
	// The string in between the tags
	Middle string
}

// A converter for BBCode
type BBConverter struct {
	mainConv *regexp2.Regexp
	codeConv *regexp2.Regexp
}

// Create a new BBConverter
func NewBBConverter() BBConverter {
	return BBConverter{
		mainConv: regexp2.MustCompile(BBMatchRegEx, regexp2.Multiline),
		codeConv: regexp2.MustCompile(BBCodeRegEx, regexp2.Multiline),
	}
}

// Converts BBCode into HTML.
func (b BBConverter) HTMLConvert(in string) string {
	var codeBlocks []string
	var match *regexp2.Match
	var err error
	// First find code blocks so we don't accidentally format it's contents
	for {
		match, err = b.codeConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + codePlaceholder + in[match.Index+match.Length:]
		codeBlocks = append(codeBlocks, match.GroupByNumber(1).String())
	}
	for {
		match, err = b.mainConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + convertMatchedTag(match) + in[match.Index+match.Length:]
	}
	for i := range codeBlocks {
		if strings.Contains(codeBlocks[i], "\n") {
			in = strings.Replace(in, codePlaceholder, "<pre><code>"+codeBlocks[i]+"</code></pre>", 1)
		} else {
			in = strings.Replace(in, codePlaceholder, "<code>"+codeBlocks[i]+"</code>", 1)
		}
	}
	return in
}

// Parse and Convert BBCode. The BBCode is replaced with the return from the given conversion function.
// The key in the map is the BBCode's tag
func (b BBConverter) CustomConvert(in string, convert map[string]func(BBTag) string) string {
	var match *regexp2.Match
	var err error
	for {
		match, err = b.mainConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + convert[match.GroupByNumber(1).String()](matchToTag(match)) + in[match.Index+match.Length:]
	}
	return in
}

func matchToTag(match *regexp2.Match) BBTag {
	out := BBTag{
		Tag:    match.GroupByNumber(1).String(),
		Middle: match.GroupByNumber(3).String(),
		Values: make(map[string]string),
	}
	params := match.GroupByNumber(2).String()
	if params == "" {
		return out
	}
	if strings.HasPrefix(params, "=") {
		if params[1] == '"' {
			endInd := strings.Index(params[2:], `"`) + 2
			if endInd != 1 {
				out.Leading = params[2:endInd]
				params = params[endInd+1:]
			}
		} else {
			endInd := strings.Index(params, " ")
			if endInd == -1 {
				out.Leading = params[1:]
				return out
			}
			out.Leading = params[1:endInd]
			params = params[endInd+1:]
		}
	}
	params = strings.TrimSpace(params)
	for params != "" {
		//TODO
	}
	return out
}

func convertMatchedTag(match *regexp2.Match) string {
	tag := match.GroupByNumber(1).String()
	params := match.GroupByNumber(2).String()
	middle := match.GroupByNumber(3).String()
	switch tag {
	case "b":
		fallthrough
	case "bold":
		fallthrough
	case "i":
		fallthrough
	case "italics":
		fallthrough
	case "s":
		fallthrough
	case "strike":
		return "<" + string(tag[0]) + ">" + middle + "</" + string(tag[0]) + ">"
	case "font":
		return "TODO"
	case "size":
		return "TODO"
	case "color":
		return "TODO"
	case "smallcaps":
		return "<span style='font-variant:small-caps'>" + middle + "</span>"
	case "url":
		fallthrough
	case "link":
		return "TODO"
	case "youtube":
		return "TODO"
	case "img":
		fallthrough
	case "image":
		return "TODO"
	case "title":
		tag = "t1"
		fallthrough
	case "t1":
		fallthrough
	case "t2":
		fallthrough
	case "t3":
		fallthrough
	case "t4":
		fallthrough
	case "t5":
		fallthrough
	case "t6":
		return "<h" + string(tag[1]) + ">" + middle + "</h" + string(tag[1]) + ">"
	case "align":
		if params == "" || !strings.HasPrefix(params, "=") {
			return middle
		}
		return "<div style='text-align:" + strings.TrimPrefix(params, "=") + "'>" + middle + "</div>"
	case "bullet":
		fallthrough
	case "ul":
		return "TODO"
	case "number":
		fallthrough
	case "ol":
		return "TODO"
	}
	return middle
}
