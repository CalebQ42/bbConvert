//Package bbConvert is a simple way to convert lines of text that contain BBCode to HTML
package bbConvert

import (
	"strings"
)

const (
	left  = "left"
	right = "right"
)

var (
	classes   string
	styleness string
)

//Convert takes in a string slice (with bbcode) and converts it to proper HTML as a single string
//If pWrap == true then each part of the slice is surrounded in paragraph tags
//If pWrap == true and it finds a list, it will wrap the list in paragraph tags
func Convert(strs []string, pWrap bool) string {
	var parsedStrs []string
	for i := 0; i < len(strs); i++ {
		v := strs[i]
		if strings.Contains(v, "[ul]") {
			for j := i; j < len(strs); j++ {
				tm := strs[j]
				var tmp string
				if strings.Contains(tm, "[/ul]") {
					for _, val := range strs[i : j+1] {
						tmp += val
					}
					parsedStrs = append(parsedStrs, tmp)
					i = j
					break
				}
			}
		} else if strings.Contains(v, "[ol]") {
			for j := i; j < len(strs); j++ {
				tm := strs[j]
				var tmp string
				if strings.Contains(tm, "[/ol]") {
					for _, val := range strs[i : j+1] {
						tmp += val
					}
					parsedStrs = append(parsedStrs, tmp)
					i = j
					break
				}
			}
		} else {
			parsedStrs = append(parsedStrs, v)
		}
	}
	var out string
	for _, v := range parsedStrs {
		var tmp string
		if pWrap {
			if styleness != "" {
				tmp += " style='" + styleness + "'"
			}
			if classes != "" {
				tmp += " class='" + classes + "'"
			}
			tmp = "<p" + tmp + ">"
			tmp += bbConv(v) + "</p>"
		} else {
			tmp = bbConv(v)
		}
		out += tmp
	}
	return out
}

/*AddParagraphClass is used to add classes to the paragraph tags wraped around the output*/
/*Can be called multiple times to add multiple classes*/
func AddParagraphClass(class string) {
	class = strings.TrimSpace(class)
	classes += " " + class
	classes = strings.TrimSpace(classes)
}

/*ClearParagraphClasses clears any set classes for paragraph tags wraped around output*/
func ClearParagraphClasses() {
	classes = ""
}

/*ClearParagraphStyle clears any set styles for paragraph tags wraped around output*/
func ClearParagraphStyle() {
	styleness = ""
}

/*AddStyle is used to add style to the paragraph tags wraped around the output*/
/*Can be called muliple times to add multiple styles*/
func AddStyle(style string) {
	style = strings.TrimSpace(style)
	if strings.HasSuffix(style, ";") {
		styleness += style
	} else {
		styleness += style + ";"
	}
}

func bbConv(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			for j := i; j < len(str); j++ {
				if str[j] == ']' {
					tmp := toHTML(str[i : j+1])
					if tmp != str[i:j+1] {
						str = str[:i] + tmp + str[j+1:]
					}
				}
			}
		}
	}
	return str
}

func toHTML(str string) string {
	var beg, end string
	for i, v := range str {
		if v == ']' || v == ' ' || v == '=' {
			beg = str[1:i]
			break
		}
	}
	var tmp string
	for i := len(str) - 3; i > 0; i-- {
		if str[i:i+2] == "[/" {
			tmp = str[i:]
			end = str[i+2 : len(str)-1]
			break
		}
	}
	if strings.ToLower(beg) != strings.ToLower(end) {
		return str
	}
	for i, v := range str {
		if v == ']' {
			beg = str[1:i]
			break
		}
	}
	if strings.HasPrefix(tmp, "[/") && strings.HasSuffix(tmp, "]") && !isBBTag(tmp[2:len(tmp)-1]) {
		return str
	}
	str = bbToTag(str[len(beg)+2:len(str)-len(tmp)], beg)
	return str
}

func isBBTag(str string) bool {
	str = strings.ToLower(str)
	tf := str == "b" || str == "i" || str == "u" || str == "s" || str == "url" || str == "img" || str == "quote" || str == "style" || str == "color" || str == "youtube" || str == "ol" || str == "ul" || str == "title"
	return tf
}

func bbToTag(in, bb string) string {
	lwrbb := strings.ToLower(bb)
	var str string
	if lwrbb == "img" {
		str = "<img style='float:left;width:20%;' src='" + in + "'/>"
	} else if strings.HasPrefix(lwrbb, "img") {
		tagness := ""
		style := make(map[string]string)
		style["float"] = left
		other := make(map[string]string)
		pos := make(map[string]int)
		if strings.Contains(lwrbb, "alt=\"") || strings.Contains(lwrbb, "alt='") {
			pos["alt"] = strings.Index(lwrbb, "alt=")
			for i := strings.Index(bb, "alt=") + 5; i < len(bb); i++ {
				if (bb[i] == bb[strings.Index(lwrbb, "alt=")+4] && bb[i-1] != '\\') || bb[i] == ']' {
					other["alt"] = bb[strings.Index(lwrbb, "alt=")+5 : i]
					pos["altEnd"] = i
					break
				}
			}
		}
		if strings.Contains(lwrbb, "title=\"") || strings.Contains(lwrbb, "title='") {
			pos["title"] = strings.Index(lwrbb, "title=")
			for i := strings.Index(lwrbb, "title=") + 7; i < len(bb); i++ {
				if (bb[i] == bb[strings.Index(lwrbb, "title=")+6] && bb[i-1] != '\\') || bb[i] == ']' {
					other["title"] = bb[strings.Index(lwrbb, "title=")+7 : i]
					pos["titleEnd"] = i
					break
				}
			}
		}
		if strings.Contains(lwrbb, "height=") {
			if (pos["alt"] == 0 || strings.Index(lwrbb, "height=") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "height=") < pos["title"]) {
				var sz string
				for i := strings.Index(lwrbb, "height=") + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[strings.Index(lwrbb, "height=")+7 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[strings.Index(lwrbb, "height=")+7 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["height"] = sz
			} else if (pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "height=") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "height=") > pos["titleEnd"]) {
				var sz string
				for i := strings.LastIndex(lwrbb, "height=") + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[strings.LastIndex(lwrbb, "height=")+7 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[strings.LastIndex(lwrbb, "height=")+7 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["height"] = sz
			}
		}
		if strings.Contains(bb, "width=") {
			if (pos["alt"] == 0 || strings.Index(lwrbb, "width=") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "width=") < pos["title"]) {
				var sz string
				for i := strings.Index(lwrbb, "width=") + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[strings.Index(lwrbb, "width=")+6 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[strings.Index(lwrbb, "width=")+6 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["width"] = sz
			} else if (pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "width=") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "width=") > pos["titleEnd"]) {
				var sz string
				for i := strings.LastIndex(lwrbb, "width=") + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[strings.LastIndex(lwrbb, "width=")+6 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[strings.LastIndex(lwrbb, "width=")+6 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["width"] = sz
			}
		}
		if strings.HasPrefix(lwrbb, "img=") {
			var sz string
			for i := 5; i < len(bb); i++ {
				if bb[i] == ' ' {
					sz = lwrbb[4:i]
				} else if i == len(bb)-1 {
					sz = lwrbb[4 : i+1]
				}
			}
			w, h := sz[:strings.Index(sz, "x")], sz[strings.Index(sz, "x")+1:]
			style["height"] = h
			style["width"] = w
		}
		if strings.Contains(lwrbb, "left") {
			if ((pos["alt"] == 0 || strings.Index(lwrbb, "left") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "left") < pos["title"])) || ((pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "left") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "left") > pos["titleEnd"])) {
				style["float"] = left
			}
		} else if strings.Contains(lwrbb, "right") {
			if ((pos["alt"] == 0 || strings.Index(lwrbb, "right") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "right") < pos["title"])) || ((pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["titleEnd"])) {
				style["float"] = right
			}
		}
		if style["height"] == "" && style["width"] == "" {
			style["width"] = "20%"
		}
		tagness = " style='"
		for i, v := range style {
			tagness += i + ":" + v + ";"
		}
		tagness += "'"
		if other["alt"] != "" {
			tagness += " alt='" + other["alt"] + "'"
		}
		if other["title"] != "" {
			tagness += " title='" + other["title"] + "'"
		}
		str = "<img" + tagness + " src='" + in + "'/>"
	} else if lwrbb == "b" || lwrbb == "i" || lwrbb == "u" || lwrbb == "s" {
		str = "<" + bb + ">" + in + "</" + bb + ">"
	} else if lwrbb == "url" {
		str = "<a href='" + str[5:len(str)-6] + "'>" + in + "</a>"
	} else if strings.HasPrefix(lwrbb, "url=") {
		str = "<a href='" + bb[5:] + "'>" + in + "</a>"
	} else if strings.HasPrefix(lwrbb, "color=") {
		str = "<span style='color:" + bb[7:] + ";'>" + in + "</span>"
	} else if strings.HasPrefix(lwrbb, "quote=\"") || strings.HasPrefix(lwrbb, "quote='") {
		str = "<div class='quote'>" + bb[7:len(bb)-1] + "<blockquote>" + in + "</blockquote></div>"
	} else if strings.HasPrefix(lwrbb, "quote=") {
		str = "<div class='quote'>" + bb[6:] + "<blockquote>" + in + "</blockquote></div>"
	} else if lwrbb == "youtube" {
		lwrin := strings.ToLower(in)
		parsed := ""
		if strings.HasPrefix(lwrin, "http://") || strings.HasPrefix(lwrin, "https://") || strings.HasPrefix(in, "youtu") || strings.HasPrefix(lwrin, "www.") {
			tmp := in[7:]
			tmp = strings.Trim(tmp, "/")
			ytb := strings.Split(tmp, "/")
			if strings.HasPrefix(ytb[len(ytb)-1], "watch?v=") {
				parsed = ytb[len(ytb)-1][8:]
			} else {
				parsed = ytb[len(ytb)-1]
			}
		} else {
			parsed = in
		}
		str = "<iframe height='315' width='560' src='https://www.youtube.com/embed/" + parsed + "' frameborder='0' allowfullscreen></iframe>"
	} else if strings.HasPrefix(bb, "youtube") {
		style := make(map[string]string)
		style["height"] = "315"
		style["width"] = "560"
		if strings.Contains(lwrbb, "height=") {
			var sz string
			for i := strings.Index(lwrbb, "height=") + 7; i < len(bb); i++ {
				if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
					sz = bb[strings.Index(lwrbb, "height=")+7 : i]
					break
				} else if i == len(bb)-1 {
					sz = bb[strings.Index(lwrbb, "height=")+7 : i+1]
					break
				}
			}
			sz = strings.Replace(sz, "\"", "", -1)
			sz = strings.Replace(sz, "'", "", -1)
			style["height"] = sz
			style["width"] = ""
		}
		if strings.Contains(lwrbb, "width=") {
			var sz string
			for i := strings.Index(lwrbb, "width=") + 7; i < len(bb); i++ {
				if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
					sz = bb[strings.Index(lwrbb, "width=")+6 : i]
					break
				} else if i == len(bb)-1 {
					sz = bb[strings.Index(lwrbb, "width=")+6 : i+1]
					break
				}
			}
			sz = strings.Replace(sz, "\"", "", -1)
			sz = strings.Replace(sz, "'", "", -1)
			style["width"] = sz
		}
		if style["height"] == "" && style["width"] == "" {
			style["height"] = "315"
			style["width"] = "560"
		}
		if strings.Contains(lwrbb, "left") {
			style["float"] = left
		}
		if strings.Contains(lwrbb, "right") {
			style["float"] = right
		}
		if strings.HasPrefix(lwrbb, "youtube=") {
			var sz string
			for i := 9; i < len(bb); i++ {
				if bb[i] == ' ' {
					sz = lwrbb[8:i]
				} else if i == len(bb)-1 {
					sz = lwrbb[8 : i+1]
				}
			}
			w, h := sz[:strings.Index(sz, "x")], sz[strings.Index(sz, "x")+1:]
			style["height"] = h
			style["width"] = w
		}
		lwrin := strings.ToLower(in)
		parsed := ""
		if strings.HasPrefix(lwrin, "http://") || strings.HasPrefix(lwrin, "https://") || strings.HasPrefix(in, "youtu") || strings.HasPrefix(lwrin, "www.") {
			tmp := in[7:]
			tmp = strings.Trim(tmp, "/")
			ytb := strings.Split(tmp, "/")
			if strings.HasPrefix(ytb[len(ytb)-1], "watch?v=") {
				parsed = ytb[len(ytb)-1][8:]
			} else {
				parsed = ytb[len(ytb)-1]
			}
		} else {
			parsed = in
		}
		str = "<iframe style='"
		for i, v := range style {
			str += i + ":" + v + ";"
		}
		str += "' src='https://www.youtube.com/embed/" + parsed + "' frameborder='0' allowfullscreen></iframe>"
	} else if lwrbb == "ul" {
		split := strings.Split(in, "*")
		for i := range split {
			split[i] = strings.TrimSpace(split[i])
		}
		for _, v := range split {
			if v != "" && v != "\n" {
				str += "<li>" + v + "</li>"
			}
		}
		str = "<ul>" + str + "</ul>"
	} else if lwrbb == "ol" {
		split := strings.Split(in, "*")
		for i := range split {
			split[i] = strings.TrimSpace(split[i])
		}
		for _, v := range split {
			if v != "" && v != "\n" {
				str += "<li>" + v + "</li>"
			}
		}
		str = "<ol>" + str + "</ol>"
	} else if lwrbb == "title" {
	}
	return str
}
