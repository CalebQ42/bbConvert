package bbConvert

import (
	"strconv"
	"strings"
)

var (
	tagFuncs map[string]func(Tag, string) string
)

func init() {
	tagFuncs = make(map[string]func(Tag, string) string)
}

//AddCustomTag allows you to add support for a custom bbCode tag. If ImplementDefaults and SetWrap(true) isn't used
//you could theoretically use this to convert to a different language besides HTML.
//If you ImplementDefaults, do it before SetConvert because it'll overwrite any conflicting custom tags.
//If you ImplementDefaults then AddCustomTag for a type already supported by the defaults, the AddCustomTag will overwrite
//the default.
func AddCustomTag(tag string, f func(Tag, string) string) {
	tagFuncs[tag] = f
}

//ImplementDefaults adds support to convert the default bbCode tags. List can be found in the README.md
func ImplementDefaults() {
	tmp := func(_ Tag, meat string) string {
		return "<b>" + meat + "</b>"
	}
	tagFuncs["b"] = tmp
	tagFuncs["bold"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<i>" + meat + "<i>"
	}
	tagFuncs["i"] = tmp
	tagFuncs["italics"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<s>" + meat + "</s>"
	}
	tagFuncs["s"] = tmp
	tagFuncs["strike"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<u>" + meat + "</u>"
	}
	tagFuncs["u"] = tmp
	tagFuncs["underline"] = tmp
	tagFuncs["font"] = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["color"] = fnt.FindValue("color")
		style["font-family"] = fnt.FindValue("starting")
		style["font-size"] = fnt.FindValue("size")
		tmp := fnt.FindValue("variant")
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
		return out
	}
	tagFuncs["color"] = func(fnt Tag, meat string) string {
		return "<span style='color:" + fnt.FindValue("starting") + ";'>" + meat + "</span>"
	}
	tagFuncs["size"] = func(fnt Tag, meat string) string {
		return "<span style='font-color:" + fnt.FindValue("starting") + ";'>" + meat + "</span>"
	}
	tagFuncs["smallcaps"] = func(_ Tag, meat string) string {
		return "<span style=\"font-variant:small-caps;\">" + meat + "</span>"
	}
	tmp = func(fnt Tag, meat string) string {
		if fnt.FindValue("starting") == "" {
			return "<a href=\"" + meat + "\">" + meat + "</a>"
		}
		return "<a href=\"" + fnt.FindValue("starting") + "\">" + meat + "</a>"
	}
	tagFuncs["url"] = tmp
	tagFuncs["link"] = tmp
	tmp = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.FindValue("right") != "" {
			style["float"] = fnt.FindValue("right")
		} else if fnt.FindValue("left") != "" {
			style["float"] = fnt.FindValue("left")
		}
		alt := fnt.FindValue("alt")
		title := fnt.FindValue("title")
		style["width"] = fnt.FindValue("width")
		style["height"] = fnt.FindValue("height")
		srt := fnt.FindValue("starting")
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
		return out
	}
	tagFuncs["img"] = tmp
	tagFuncs["image"] = tmp
	tagFuncs["youtube"] = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.FindValue("right") != "" {
			style["float"] = fnt.FindValue("right")
		} else if fnt.FindValue("left") != "" {
			style["float"] = fnt.FindValue("left")
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
		return out
	}
	tmp = func(fnt Tag, meat string) string {
		var out string
		meat = strings.Replace(strings.Replace(meat, p, "", -1), "</p>", "", -1)
		out = bulletprocessing(meat)
		out = "<ul>" + out + "</ul>"
		if pWrap {
			out = strings.Replace(out, "\n", "", -1)
			out = "</p>" + out + p
		}
		return out
	}
	tagFuncs["ul"] = tmp
	tagFuncs["bullet"] = tmp
	tmp = func(fnt Tag, meat string) string {
		var out string
		meat = strings.Replace(strings.Replace(meat, p, "", -1), "</p>", "", -1)
		out = bulletprocessing(meat)
		out = "<ol>" + out + "</ol>"
		if pWrap {
			out = strings.Replace(out, "\n", "", -1)
			out = "</p>" + out + p
		}
		return out
	}
	tagFuncs["ol"] = tmp
	tagFuncs["number"] = tmp
	tagFuncs["title"] = func(fnt Tag, meat string) string {
		var out string
		meat = strings.Replace(meat, "\n", "", -1)
		out = "<h1>" + meat + "</h1>"
		if pWrap {
			out = "</p>" + out + p
		}
		return out
	}
	tagFuncs["align"] = func(fnt Tag, meat string) string {
		var out string
		out = "<div style=\"text-align:" + fnt.FindValue("starting") + ";\">"
		if pWrap {
			out = "</p>" + out + p + meat + "</p></div>" + p
		} else {
			out = out + meat + "</div>"
		}
		return out
	}
	tagFuncs["t1"] = func(fnt Tag, meat string) string {
		return "<h1>" + meat + "</h1>"
	}
	tagFuncs["t2"] = func(fnt Tag, meat string) string {
		return "<h2>" + meat + "</h2>"
	}
	tagFuncs["t3"] = func(fnt Tag, meat string) string {
		return "<h3>" + meat + "</h3>"
	}
	tagFuncs["t4"] = func(fnt Tag, meat string) string {
		return "<h4>" + meat + "</h4>"
	}
	tagFuncs["t5"] = func(fnt Tag, meat string) string {
		return "<h5>" + meat + "</h5>"
	}
	tagFuncs["t6"] = func(fnt Tag, meat string) string {
		return "<h6>" + meat + "</h6>"
	}
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
