package bbConvert

import (
	"strconv"
	"strings"
)

func toHTML(fnt, end tag, meat string) (out string) {
	switch fnt.bbType {
	case "b", "bold":
		out = "<b>" + meat + "</b>"
	case "i", "italics":
		out = "<i>" + meat + "</i>"
	case "s", "strike":
		out = "<s>" + meat + "</s>"
	case "u", "underline":
		out = "<u>" + meat + "</u>"
	case "font":
		style := make(map[string]string)
		style["color"] = fnt.findValue("color")
		style["font-family"] = fnt.findValue("starting")
		style["font-size"] = fnt.findValue("size")
		tmp := fnt.findValue("variant")
		if tmp == "smallcaps" {
			style["font-variant"] = "small-caps"
		} else if tmp == "upper" {
			style["text-transform"] = "uppercase"
		} else if tmp == "lower" {
			style["text-transform"] = "lowercase"
		}
		if !strings.HasPrefix(style["color"], "#") {
			_, err := strconv.ParseInt(style["color"], 16, 0)
			if err == nil {
				style["color"] = "#" + style["color"]
			}
		}
		style = trimAccess(style)
		if len(style) != 0 {
			out = "<span style=\""
			for i, v := range style {
				if strings.Contains(v, " ") {
					v = "'" + v + "'"
				}
				out += i + ":" + v + ";"
			}
			out += "\">" + meat + "</span>"
		} else {
			return meat
		}
	case "color":
		out = "<span style='color:" + fnt.findValue("starting") + ";'>" + meat + "</span>"
	case "size":
		out = "<span style='font-color:" + fnt.findValue("starting") + ";'>" + meat + "</span>"
	case "smallcaps":
		out = "<span style=\"font-variant:small-caps;\">" + meat + "</span>"
	case "url", "link":
		if fnt.findValue("starting") == "" {
			out = "<a href=\"" + meat + "\">" + meat + "</a>"
		} else {
			out = "<a href=\"" + fnt.findValue("starting") + "\">" + meat + "</a>"
		}
	case "img", "image":
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.findValue("right") != "" {
			style["float"] = fnt.findValue("right")
		} else if fnt.findValue("left") != "" {
			style["float"] = fnt.findValue("left")
		}
		alt := fnt.findValue("alt")
		title := fnt.findValue("title")
		style["width"] = fnt.findValue("width")
		style["height"] = fnt.findValue("height")
		srt := fnt.findValue("starting")
		if srt != "" {
			if ind := strings.Index(srt, "x"); ind != -1 {
				style["width"] = srt[:ind]
				style["height"] = srt[ind+1:]
			} else {
				style["width"] = srt
			}
		}
		out = "<img "
		style = trimAccess(style)
		if style["width"] == "" && style["height"] == "" {
			style["width"] = "20%"
		}
		out += "style='"
		for i, v := range style {
			out += i + ":" + v + ";"
		}
		out += "' "
		if alt != "" {
			out += "alt='" + strings.Replace(alt, "'", "\\'", -1) + "' "
		}
		if title != "" {
			out += "title='" + strings.Replace(title, "'", "\\'", -1) + "' "
		}
		out += "src='" + meat + "'/>"
	case "youtube":
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.findValue("right") != "" {
			style["float"] = fnt.findValue("right")
		} else if fnt.findValue("left") != "" {
			style["float"] = fnt.findValue("left")
		}
		lwrin := strings.ToLower(meat)
		var parsed string
		if strings.HasPrefix(lwrin, "http://") || strings.HasPrefix(lwrin, "https://") || strings.HasPrefix(lwrin, "youtu") || strings.HasPrefix(lwrin, "www.") {
			tmp := meat[7:]
			tmp = strings.Trim(tmp, "/")
			ytb := strings.Split(tmp, "/")
			if strings.HasPrefix(ytb[len(ytb)-1], "watch?v=") {
				parsed = ytb[len(ytb)-1][8:]
			} else {
				parsed = ytb[len(ytb)-1]
			}
		} else {
			parsed = meat
		}
		out = "<ifram"
		if len(style) != 0 {
			out += "style='"
			for i, v := range style {
				out += i + ":" + v + ";"
			}
			out += "' "
		}
		out += " src='https://www.youtube.com/embed/" + parsed + "' frameborder='0' allowfullscreen></iframe>"
	case "ul", "bullet":
		meat = strings.Replace(strings.Replace(meat, p, "", -1), "</p>", "", -1)
		out = bulletprocessing(meat)
		out = "<ul>" + out + "</ul>"
		if pWrap {
			out = strings.Replace(out, "\n", "", -1)
			out = "</p>" + out + p
		}
	case "ol", "number":
		meat = strings.Replace(strings.Replace(meat, p, "", -1), "</p>", "", -1)
		out = bulletprocessing(meat)
		out = "<ol>" + out + "</ol>"
		if pWrap {
			out = strings.Replace(out, "\n", "", -1)
			out = "</p>" + out + p
		}
	case "title":
		meat = strings.Replace(meat, "\n", "", -1)
		out = "<h1>" + meat + "</h1>"
		if pWrap {
			out = "</p>" + out + p
		}
	case "align":
		out = "<div style=\"text-align:" + fnt.findValue("starting") + ";\">"
		if pWrap {
			out = "</p>" + out + p + meat + "</p></div>" + p
		} else {
			out = out + meat + "</div>"
		}
	default:
		if strings.HasPrefix(fnt.bbType, "t") {
			meat = strings.Replace(meat, "\n", "", -1)
			par, err := strconv.Atoi(fnt.bbType[1:])
			if err == strconv.ErrSyntax {
				out = "<h4>" + meat + "</h4>"
			} else {
				if par >= 1 && par <= 6 {
					out = "<h" + strconv.Itoa(par) + ">" + meat + "</h" + strconv.Itoa(par) + ">"
				} else if par < 1 {
					out = "<h1>" + meat + "</h1>"
				} else if par > 6 {
					out = "<h6>" + meat + "</h6>"
				}
			}
			if pWrap {
				out = "</p>" + out + p
			}
		} else {
			out = fnt.fullBB + meat + end.fullBB
		}
	}
	return
}

func trimAccess(style map[string]string) (out map[string]string) {
	out = style
	for i, v := range style {
		if v == "" {
			delete(out, i)
		}
	}
	return
}

func bulletprocessing(meat string) (out string) {
	var prev int
	var bullets []string
	for i := 0; i < len(meat); i++ {
		v := meat[i]
		if i+4 < len(meat) && (meat[i:i+4] == "<ul>" || meat[i:i+4] == "<ol>") {
			var count int
			if meat[i:i+4] == "<ul>" {
				for j := i + 4; j < len(meat)-5; j++ {
					if meat[j:j+4] == "<ul>" {
						count++
					} else if meat[j:j+5] == "</ul>" {
						count--
						if count == -1 {
							bullets = append(bullets, meat[prev:j+5])
							i = j + 5
							prev = j + 6
							break
						}
					}
				}
			} else if meat[i:i+4] == "<ol>" {
				for j := i + 4; j < len(meat)-5; j++ {
					if meat[j:j+4] == "<ol>" {
						count++
					} else if meat[j:j+5] == "</ol>" {
						count--
						if count == -1 {
							bullets = append(bullets, meat[prev:j+5])
							i = j + 5
							prev = j + 6
							break
						}
					}
				}
			}
		} else if v == '*' || v == '\n' {
			if prev != i {
				bullets = append(bullets, meat[prev:i])
			}
			prev = i + 1
		}
	}
	if meat[prev:] != "" {
		bullets = append(bullets, meat[prev:])
	}
	for _, v := range bullets {
		v = strings.TrimSpace(v)
		if (strings.HasPrefix(v, "<ul>") && strings.HasSuffix(v, "</ul>")) || (strings.HasPrefix(v, "<ol>") && strings.HasSuffix(v, "</ol>")) {
			out += v
		} else {
			if v != "" && v != "\n" {
				out += "<li>" + v + "</li>"
			}
		}
	}
	out = strings.Replace(out, "<li></li>", "", -1)
	return
}
