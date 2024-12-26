//go:generate regexp2cg -o regexp2_codegen.go
package bbConvert

import (
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

// The Magic RegEx string that matches bbCode tags.
const (
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
	codeConv   *regexp2.Regexp
	mainConv   *regexp2.Regexp
	customConv *regexp2.Regexp
}

func NewBBConverter() BBConverter {
	code := regexp2.MustCompile(`\[code\]([\s\S]*?)\[\/code\]`, regexp2.Multiline)
	main := regexp2.MustCompile(`\[(b|bold|i|italics|u|underline|s|strike|font|size|color|smallcaps|url|link|youtube|img|image|title|t[1-6]|align|float|ul|bullet|ol|number)(.*?)\]([\s\S]*?)\[\/\1\]`, regexp2.Multiline)
	custom := regexp2.MustCompile(`\[(\w+\b)(.*?)\]([\s\S]*?)\[\/\1\]`, regexp2.Multiline)
	return BBConverter{
		codeConv:   code,
		mainConv:   main,
		customConv: custom,
	}
}

// Converts BBCode into HTML.
func (b BBConverter) HTMLConvert(bb string) string {
	return b.bbActualConv([]rune(bb), false)
}

func (b BBConverter) bbActualConv(in []rune, comboConv bool) string {
	var codeBlocks []string
	var match *regexp2.Match
	var err error
	if !comboConv {
		// First find code blocks so we don't accidentally format it's contents
		for {
			match, err = b.codeConv.FindRunesMatch(in)
			if err != nil || match == nil {
				break
			}
			in = slices.Concat(in[:match.Index], []rune(codePlaceholder), in[match.Index+match.Length:])
			codeBlocks = append(codeBlocks, match.GroupByNumber(1).String())
		}
	}
	for {
		match, err = b.mainConv.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		in = slices.Concat(in[:match.Index], []rune(matchToHTML(match)), in[match.Index+match.Length:])
	}
	out := string(in)
	if !comboConv {
		out = "<p>" + strings.ReplaceAll(out, "\n", "</p>\n<p>") + "</p>"
		for i := range codeBlocks {
			if strings.Contains(codeBlocks[i], "\n") {
				out = strings.Replace(out, codePlaceholder, "<pre><code>"+codeBlocks[i]+"</code></pre>", 1)
			} else {
				out = strings.Replace(out, codePlaceholder, "<code>"+codeBlocks[i]+"</code>", 1)
			}
		}
	}
	return out

}

func matchToHTML(match *regexp2.Match) string {
	tag := match.GroupByNumber(1).String()
	leading, params := processParams(match.GroupByNumber(2).String())
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
		if leading == "" && len(params) == 0 {
			return middle
		}
		style := "style='"
		if leading != "" {
			style += "font-family:" + leading + ";"
		}
		if params["color"] != "" {
			style += "color:" + params["color"] + ";"
		}
		if params["size"] != "" {
			style += "font-size:" + params["size"] + ";"
		}
		switch params["variant"] {
		case "upper":
			style += "text-transform:uppercase;"
		case "lower":
			style += "text-transform:lowercase;"
		case "smallcaps":
			style += "font-variant:small-caps;"
		}
		return "<span " + style + "'>" + middle + "</span>"
	case "size":
		if leading == "" {
			return middle
		}
		return "<span style='font-size:" + leading + "'>" + middle + "</span>"
	case "color":
		if leading == "" {
			return middle
		}
		return "<span style='color:" + leading + "'>" + middle + "</span>"
	case "smallcaps":
		return "<span style='font-variant:small-caps'>" + middle + "</span>"
	case "url":
		fallthrough
	case "link":
		var addr, extras string
		if params["title"] != "" {
			extras = "title=\"" + params["title"] + "\""
		}
		if _, exist := params["tab"]; exist {
			extras += "target='_blank'"
		}
		if leading == "" {
			addr = middle
		} else {
			addr = leading
		}
		return "<a href='" + addr + "'" + extras + ">" + middle + "</a>"
	case "youtube":
		if strings.Contains(middle, "/") {
			ytUrl, err := url.Parse(middle)
			if err != nil {
				return middle
			}
			if ytUrl.Path == "/watch" {
				if ytUrl.Query().Get("v") != "" {
					middle = ytUrl.Query().Get("v")
				} else {
					return middle
				}
			} else {
				spl := strings.Split(ytUrl.Path, "/")
				if len(spl) == 0 {
					return middle
				}
				middle = spl[len(spl)-1]
			}
		}
		middle = "https://youtube.com/embed/" + middle
		var style string
		var width, height string
		if leading != "" {
			xInd := strings.Index(leading, "x")
			if xInd != -1 {
				width = leading[:xInd]
				height = leading[xInd+1:]
			}
		} else {
			width = params["width"]
			height = params["height"]
		}
		if width != "" {
			// Does it contain units? if not, we assume pixels (px)
			_, err := strconv.Atoi(width)
			if err == nil {
				width += "px"
			}
			style += "width:" + width + ";"
		}
		if height != "" {
			// Does it contain units? if not, we assume pixels (px)
			_, err := strconv.Atoi(height)
			if err == nil {
				height += "px"
			}
			style += "height:" + height + ";"
		}
		if _, exist := params["left"]; exist {
			style += "float:left;"
		} else if _, exist = params["right"]; exist {
			style += "float:right;"
		}
		if style == "" {
			return "<iframe src='" + middle + "' allowfullscreen></iframe>"
		}
		return "<iframe src='" + middle + "' style='" + style + "' allowfullscreen></iframe>"
	case "img":
		fallthrough
	case "image":
		out := "<img src='" + middle + "'"
		if params["alt"] != "" {
			out += " alt=\"" + params["alt"] + "\""
		}
		if params["title"] != "" {
			out += " title=\"" + params["title"] + "\""
		}
		var style string
		var width, height string
		if leading != "" {
			xInd := strings.Index(leading, "x")
			if xInd != -1 {
				width = leading[:xInd]
				height = leading[xInd+1:]
			}
		} else {
			width = params["width"]
			height = params["height"]
		}
		if width != "" {
			// Does it contain units? if not, we assume pixels (px)
			_, err := strconv.Atoi(width)
			if err == nil {
				width += "px"
			}
			style += "width:" + width + ";"
		}
		if height != "" {
			// Does it contain units? if not, we assume pixels (px)
			_, err := strconv.Atoi(height)
			if err == nil {
				height += "px"
			}
			style += "height:" + height + ";"
		}
		if _, exist := params["left"]; exist {
			style += "float:left;"
		} else if _, exist = params["right"]; exist {
			style += "float:right;"
		}
		if style != "" {
			out += " style='" + style + "'"
		}
		return out + "/>"
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
		if leading == "" && len(params) == 0 {
			return middle
		}
		align := leading
		if leading == "" {
			for k := range params {
				align = k
				break
			}
		}
		return "<div style='text-align:" + align + ";'>" + middle + "</div>"
	case "float":
		if leading == "" && len(params) == 0 {
			return middle
		}
		float := leading
		if leading == "" {
			for k := range params {
				float = k
				break
			}
		}
		return "<div style='float:" + float + ";'>" + middle + "</div>"
	case "bullet":
		tag = "ul"
		fallthrough
	case "number":
		tag = "ol"
		fallthrough
	case "ol":
		fallthrough
	case "ul":
		return "<" + tag + ">" + processListItems(middle) + "</" + tag + ">"
	}
	return middle
}

func processListItems(in string) string {
	in = strings.TrimSpace(in)
	if in == "" {
		return ""
	}
	var out []string
	for ind := strings.IndexAny(in, "\n*"); ind != -1; ind = strings.IndexAny(in, "\n*") {
		line := strings.TrimSpace(in[:ind])
		if line != "" {
			out = append(out, strings.TrimSpace(in[:ind]))
		}
		in = in[ind+1:]
	}
	in = strings.TrimSpace(in)
	if in != "" {
		out = append(out, in)
	}
	return "<li>" + strings.Join(out, "</li><li>") + "</li>"
}

type BBConvert func(BBTag) string

// Parse and Convert BBCode. The BBCode is replaced with the return from the given conversion function.
// The key in the map is the BBCode's tag
func (b BBConverter) CustomConvert(bb string, convert map[string]BBConvert) string {
	in := []rune(bb)
	var match *regexp2.Match
	var err error
	for {
		match, err = b.customConv.FindRunesMatch(in)
		if err != nil || match == nil {
			break
		}
		in = slices.Concat(in[:match.Index], []rune(convert[match.GroupByNumber(1).String()](matchToTag(match))), in[match.Index+match.Length:])
	}
	return string(in)
}

func matchToTag(match *regexp2.Match) BBTag {
	out := BBTag{
		Tag:    match.GroupByNumber(1).String(),
		Middle: match.GroupByNumber(3).String(),
	}
	out.Leading, out.Values = processParams(match.GroupByNumber(2).String())
	return out
}

func processParams(params string) (leading string, out map[string]string) {
	out = make(map[string]string)
	params = strings.TrimSpace(params)
	if params == "" {
		return
	}
	if strings.HasPrefix(params, "=") {
		if params[1] == '"' {
			endInd := strings.Index(params[2:], "\"")
			if endInd == -1 {
				leading = params[2:]
				return
			} else {
				leading = params[2 : endInd+2]
				params = params[endInd+3:]
			}
		} else {
			endInd := strings.Index(params[1:], " ")
			if endInd == -1 {
				leading = params[1:]
				return
			} else {
				leading = params[1 : endInd+1]
				params = params[endInd+2:]
			}
		}
	}
	var ind int
	for {
		params = strings.TrimSpace(params)
		if params == "" {
			break
		}
		ind = strings.IndexAny(params, " =")
		if ind == -1 || ind == len(params)-1 {
			out[params] = ""
			break
		}
		if params[ind] == ' ' {
			out[params[:ind]] = ""
			params = params[ind+1:]
			continue
		}
		key := params[:ind]
		var endInd int
		if params[ind+1] == '"' {
			endInd = strings.IndexByte(params[ind+2:], '"')
			if endInd == -1 {
				endInd = len(params)
			} else {
				endInd += ind + 2
			}
			ind += 2
		} else {
			endInd = strings.IndexByte(params[ind+1:], ' ')
			if endInd == -1 {
				endInd = len(params)
			} else {
				endInd += ind + 1
			}
			ind += 1
		}
		out[key] = params[ind:endInd]
		for endInd < len(params) && (params[endInd] == ' ' || params[endInd] == '"') {
			endInd++
		}
		params = params[endInd:]
	}
	return
}
