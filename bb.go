package bbConvert

import (
	"strings"

	"github.com/dlclark/regexp2"
)

// The Magic RegEx string that matches bbCode tags.
const BBMatchRegEx = `\[(b|bold|i|italics|u|underline|s|strike|font|size|color|smallcaps|url|link|youtube|img|image|title|t[1-6]|align|float|ul|bullet|ol|number)(.*?)\]([\s\S]*?)\[\/\1\]`

type BBConverter struct {
	conv *regexp2.Regexp
}

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

func NewBBConverter() BBConverter {
	return BBConverter{
		conv: regexp2.MustCompile(BBMatchRegEx, regexp2.Multiline),
	}
}

func (b BBConverter) HTMLConvert(in string) string {
	inRune := []rune(in)
	var match *regexp2.Match
	var err error
	for {
		match, err = b.conv.FindRunesMatch(inRune)
		if err != nil || match == nil {
			break
		}
		inRune = append(inRune[:match.Index], append(convertMatchedTag(match), inRune[match.Index+match.Length:]...)...)
	}
	return string(inRune)
}

func (b BBConverter) CustomConvert(in string, convert map[string]func(BBTag)) {

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

	}
	return out
}

func convertMatchedTag(match *regexp2.Match) []rune {
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
		return []rune(
			"<" + string(tag[0]) + ">" + middle + "</" + string(tag[0]) + ">",
		)
	case "font":
		return []rune("TODO")
	case "size":
		return []rune("TODO")
	case "color":
		return []rune("TODO")
	case "smallcaps":
		return []rune(
			"<span style='font-variant:small-caps'>" + middle + "</span>",
		)
	case "url":
		fallthrough
	case "link":
		return []rune("TODO")
	case "youtube":
		return []rune("TODO")
	case "img":
		fallthrough
	case "image":
		return []rune("TODO")
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
		return []rune(
			"<h" + string(tag[1]) + ">" + middle + "</h" + string(tag[1]) + ">",
		)
	case "align":
		if params == "" || !strings.HasPrefix(params, "=") {
			return []rune(middle)
		}
		return []rune(
			"<div style='text-align:" + strings.TrimPrefix(params, "=") + "'>" + middle + "</div>",
		)
	case "bullet":
		fallthrough
	case "ul":
		return []rune("TODO")
	case "number":
		fallthrough
	case "ol":
		return []rune("TODO")
	}
	return []rune("tested")
}
