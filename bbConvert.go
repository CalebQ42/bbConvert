//Package bbConvert is a simple way to convert lines of text that contain BBCode to HTML
package bbConvert

import (
	"strconv"
	"strings"
)

const (
	left      = "left"
	right     = "right"
	smallcaps = "smallcaps"
)

var (
	classes   string
	styleness string
	paraWrap  bool
)

//Convert takes in a string slice (with bbcode) and converts it to proper HTML as a single string
//If pWrap == true then each part of the slice is surrounded in paragraph tags
//If pWrap == true and it finds a list, it will wrap the list in paragraph tags
func Convert(strs []string, pWrap bool) string {
	paraWrap = pWrap
	var tmp []string
	for _, v := range strs {
		split := strings.Split(v, "\n")
		for _, v := range split {
			tmp = append(tmp, v)
		}
	}
	strs = tmp
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
			in := bbConv(v)
			if strings.HasPrefix(in, "<h") {
				tmp = in
			} else if v != "" {
				if styleness != "" {
					tmp += " style='" + styleness + "'"
				}
				if classes != "" {
					tmp += " class='" + classes + "'"
				}
				tmp = "<p" + tmp + ">"
				tmp += in + "</p>"
			}
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
	tf := str == "b" || str == "i" || str == "u" || str == "s" || str == "url" || str == "img" || str == "quote" || str == "style" || str == "color" || str == "youtube" || str == "ol" || str == "ul" || str == "title" || strings.HasPrefix(str, "t") || str == "font" || str == "size" || str == "smallcaps"
	return tf
}

func bbToTag(in, bb string) string {
	bb = strings.TrimSpace(bb)
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
			for i := pos["alt"] + 5; i < len(bb); i++ {
				if (bb[i] == bb[pos["alt"]+4] && bb[i-1] != '\\') || i == len(bb)-1 {
					other["alt"] = bb[pos["alt"]+5 : i]
					pos["altEnd"] = i
					break
				}
			}
		}
		if strings.Contains(lwrbb, "title=\"") || strings.Contains(lwrbb, "title='") {
			pos["title"] = strings.Index(lwrbb, "title=")
			for i := pos["title"] + 7; i < len(bb); i++ {
				if (bb[i] == bb[pos["title"]+6] && bb[i-1] != '\\') || i == len(bb)-1 {
					other["title"] = bb[pos["title"]+7 : i]
					pos["titleEnd"] = i
					break
				}
			}
		}
		if strings.Contains(lwrbb, "height=") {
			pos["height"] = strings.Index(lwrbb, "height")
			pos["lheight"] = strings.LastIndex(lwrbb, "height")
			if (pos["alt"] == 0 || pos["height"] < pos["alt"]) && (pos["title"] == 0 || pos["height"] < pos["title"]) {
				var sz string
				for i := pos["height"] + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[pos["height"]+7 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[pos["height"]+7 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["height"] = sz
			} else if (pos["altEnd"] == 0 || pos["lheight"] > pos["altEnd"]) && (pos["titleEnd"] == 0 || pos["lheight"] > pos["titleEnd"]) {
				var sz string
				for i := pos["lheight"] + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[pos["lheight"]+7 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[pos["lheight"]+7 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["height"] = sz
			} else {
				he := 0
				cnt := strings.Count(lwrbb, "height=")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "height=")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						var sz string
						for i := he + 7; i < len(bb); i++ {
							if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
								sz = bb[he+7 : i]
								break
							} else if i == len(bb)-1 {
								sz = bb[he+7 : i+1]
								break
							}
						}
						sz = strings.Replace(sz, "\"", "", -1)
						sz = strings.Replace(sz, "'", "", -1)
						style["height"] = sz
						break
					} else {
						tmp = tmp[pos["height"]+1:]
					}
				}
			}
		}
		if strings.Contains(bb, "width=") {
			pos["width"] = strings.Index(lwrbb, "width=")
			pos["lwidth"] = strings.LastIndex(lwrbb, "width=")
			if (pos["alt"] == 0 || pos["width"] < pos["alt"]) && (pos["title"] == 0 || pos["width"] < pos["title"]) {
				var sz string
				for i := pos["width"] + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[pos["width"]+6 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[pos["width"]+6 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["width"] = sz
			} else if (pos["altEnd"] == 0 || pos["lwidth"] > pos["altEnd"]) && (pos["titleEnd"] == 0 || pos["lwidth"] > pos["titleEnd"]) {
				var sz string
				for i := pos["lwidth"] + 7; i < len(bb); i++ {
					if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
						sz = bb[pos["lwidth"]+6 : i]
						break
					} else if i == len(bb)-1 {
						sz = bb[pos["lwidth"]+6 : i+1]
						break
					}
				}
				sz = strings.Replace(sz, "\"", "", -1)
				sz = strings.Replace(sz, "'", "", -1)
				style["width"] = sz
			} else {
				he := 0
				cnt := strings.Count(lwrbb, "width=")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "width=")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						var sz string
						for i := he + 6; i < len(bb); i++ {
							if bb[i] == ' ' || bb[i] == '"' || bb[i] == '\'' {
								sz = bb[he+6 : i]
								break
							} else if i == len(bb)-1 {
								sz = bb[he+6 : i+1]
								break
							}
						}
						sz = strings.Replace(sz, "\"", "", -1)
						sz = strings.Replace(sz, "'", "", -1)
						style["width"] = sz
						break
					} else {
						tmp = tmp[pos["width"]+1:]
					}
				}
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
			xPos := strings.Index(sz, "x")
			w, h := "", ""
			if xPos != -1 {
				w, h = sz[:xPos], sz[xPos+1:]
			}

			style["height"] = h
			style["width"] = w
		}
		if strings.Contains(lwrbb, "left") {
			if ((pos["alt"] == 0 || strings.Index(lwrbb, "left") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "left") < pos["title"])) || ((pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "left") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "left") > pos["titleEnd"])) {
				style["float"] = left
			} else {
				he := 0
				pos["float"] = strings.Index(lwrbb, "left")
				cnt := strings.Count(lwrbb, "left")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "left")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						style["float"] = left
						break
					} else {
						tmp = tmp[he+1:]
					}
				}
			}
		} else if strings.Contains(lwrbb, "right") {
			if ((pos["alt"] == 0 || strings.Index(lwrbb, "right") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "right") < pos["title"])) || ((pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["titleEnd"])) {
				style["float"] = right
			} else {
				he := 0
				pos["float"] = strings.Index(lwrbb, "right")
				cnt := strings.Count(lwrbb, "right")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "right")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						style["float"] = "right"
						break
					} else {
						tmp = tmp[pos["float"]+1:]
					}
				}
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
	} else if strings.HasPrefix(lwrbb, "url") {
		var url string
		if strings.HasPrefix(lwrbb, "url=") {
			for i := 4; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' || i == len(lwrbb)-1 {
					url = bb[4:i]
					break
				}
			}
		}
		var title string
		titlePos := strings.Index(lwrbb, "title=")
		if strings.Contains(lwrbb, "title='") || strings.Contains(lwrbb, "title=\"") {
			for i := titlePos + 7; i < len(bb); i++ {
				if (bb[i] == bb[titlePos+6] && bb[i-1] != '\\') || i == len(bb)-1 {
					title = bb[titlePos+7 : i]
					break
				}
			}
		}
		str = "<a"
		if title != "" {
			str += " title='" + title + "'"
		}
		if url == "" {
			url = in
		}
		str += " href='" + url + "'" + ">" + in + "</a>"
	} else if strings.HasPrefix(lwrbb, "color=") {
		tmp := bb[6:]
		if !strings.HasPrefix(tmp, "#") {
			tmp = "#" + tmp
		}
		_, err := strconv.ParseInt(tmp, 16, 0)
		if err != strconv.ErrSyntax {
			tmp = tmp[1:]
		}
		str = "<span style='color:" + tmp + ";'>" + in + "</span>"
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
		if paraWrap {
			str = "</p>" + str + "<p>"
		}
	} else if lwrbb == "title" {
		str = "<h1>" + in + "</h1>"
	} else if strings.HasPrefix(lwrbb, "t") {
		out, err := strconv.Atoi(lwrbb[1:])
		if err == strconv.ErrSyntax {
			str = in
		} else {
			if out >= 1 && out <= 6 {
				str = "<h" + strconv.Itoa(out) + ">" + in + "</h" + strconv.Itoa(out) + ">"
			} else {
				if out < 1 {
					str = "<h1>" + in + "</h1>"
				} else if out > 6 {
					str = "<h6>" + in + "</h6>"
				}
			}
		}
	} else if strings.HasPrefix(lwrbb, "font") {
		style := make(map[string]string)
		if strings.HasPrefix(lwrbb, "font=") {
			if bb[5] == '"' || bb[5] == '\'' {
				for i := 6; i < len(bb); i++ {
					if bb[i] == bb[5] {
						style["font-family"] = lwrbb[6:i]
						style["font-family"] = strings.Replace(style["font-family"], "'", "", -1)
						style["font-family"] = strings.Replace(style["font-family"], "\"", "", -1)
						break
					}
				}
			} else {
				for i := 5; i < len(bb); i++ {
					if bb[i] == ' ' || i == len(bb)-1 {
						style["font-family"] = lwrbb[6:i]
						style["font-family"] = strings.Replace(style["font-family"], "'", "", -1)
						style["font-family"] = strings.Replace(style["font-family"], "\"", "", -1)
						style["font-family"] = strings.TrimSpace(style["font-family"])
						break
					}
				}
			}
		}
		if strings.Contains(lwrbb, "size=") {
			szPos := strings.Index(lwrbb, "size=")
			for i := szPos + 5; i < len(bb); i++ {
				if bb[i] == ' ' || i == len(bb)-1 {
					style["font-size"] = lwrbb[szPos+5 : i+1]
					style["font-size"] = strings.TrimSpace(style["font-size"])
					break
				}
			}
		}
		if strings.Contains(lwrbb, "color=") {
			clPos := strings.Index(lwrbb, "color=")
			for i := clPos + 6; i < len(bb); i++ {
				if bb[i] == ' ' || i == len(bb)-1 {
					style["color"] = lwrbb[clPos+6 : i+1]
					style["color"] = strings.TrimSpace(style["color"])
					if strings.HasPrefix(style["color"], "#") {
						style["color"] = "#" + style["color"]
					}
					_, err := strconv.ParseInt(style["color"], 16, 0)
					if err == strconv.ErrSyntax {
						style["color"] = style["color"][1:]
					}
					break
				}
			}
		}
		if strings.Contains(lwrbb, "variant=") {
			vrPos := strings.Index(lwrbb, "variant=")
			for i := vrPos + 8; i < len(bb); i++ {
				if bb[i] == ' ' || i == len(bb)-1 {
					vari := lwrbb[vrPos+8 : i+1]
					vari = strings.TrimSpace(vari)
					if vari == "lower" {
						in = strings.ToLower(in)
					} else if vari == "upper" {
						in = strings.ToUpper(in)
					} else if vari == smallcaps {
						style["font-variant"] = "small-caps"
					}
				}
			}
		}
		str = "<span style=\""
		for i, v := range style {
			if strings.Contains(v, " ") {
				v = "'" + v + "'"
			}
			str += i + ":" + v + ";"
		}
		str += "\">" + in + "</span>"
	} else if strings.HasPrefix(lwrbb, "size=") {
		sz := lwrbb[5:]
		str = "<span style='font-size:" + sz + ";'>" + in + "</span>"
	} else if lwrbb == smallcaps {
		str = "<span style='font-variant:small-caps;'>" + in + "</span>"
	} else {
		return in
	}
	return str
}
