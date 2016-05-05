package bbConvert

import (
	"strconv"
	"strings"
)

const (
	lf = "left"
	rt = "right"
)

func toHTML(bb, meat string) string {
	bb = strings.TrimSpace(bb)
	lwrbb := strings.ToLower(bb)
	var str string
	if lwrbb == "b" || lwrbb == "s" || lwrbb == "i" || lwrbb == "u" {
		str = "<" + lwrbb + ">" + meat + "</" + lwrbb + ">"
	} else if lwrbb == "bold" {
		str = "<b>" + meat + "</b>"
	} else if lwrbb == "italics" {
		str = "<i>" + meat + "</i>"
	} else if lwrbb == "underline" {
		str = "<u>" + meat + "</u>"
	} else if lwrbb == "strike" {
		str = "<s>" + meat + "</s>"
	} else if strings.HasPrefix(lwrbb, "font") {
		style := make(map[string]string)
		if strings.HasPrefix(lwrbb, "font=") {
			if lwrbb[5] == '\'' || lwrbb[5] == '"' {
				for i := 6; i < len(lwrbb); i++ {
					if lwrbb[i] == lwrbb[5] {
						style["font-family"] = lwrbb[6:i]
						break
					}
				}
			} else {
				for i := 5; i < len(lwrbb); i++ {
					if lwrbb[i] == ' ' {
						style["font-family"] = lwrbb[5:i]
					}
				}
				if style["font-family"] == "" {
					style["font-family"] = lwrbb[5:]
				}
			}
		}
		if strings.Contains(lwrbb, "color=") {
			pos := strings.Index(lwrbb, "color=") + 6
			for i := pos; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' {
					style["color"] = lwrbb[pos:i]
				}
			}
			if style["color"] == "" {
				style["color"] = lwrbb[pos:]
			}
			if !strings.HasPrefix(style["color"], "#") {
				_, err := strconv.ParseInt(style["color"], 16, 0)
				if err == nil {
					style["color"] = "#" + style["color"]
				}
			}
		}
		if strings.Contains(lwrbb, "size=") {
			pos := strings.Index(lwrbb, "size=") + 5
			for i := pos; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' {
					style["font-size"] = lwrbb[pos:i]
				}
			}
			if style["font-size"] == "" {
				style["font-size"] = lwrbb[pos:]
			}
		}
		if strings.Contains(lwrbb, "variant=") {
			pos := strings.Index(lwrbb, "variant=") + 8
			var tmp string
			for i := pos; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' {
					tmp = lwrbb[pos:i]
				}
			}
			if tmp == "" {
				tmp = lwrbb[pos:]
			}
			if tmp == "smallcaps" {
				style["font-variant"] = "small-caps"
			} else if tmp == "upper" {
				style["text-transform"] = "uppercase"
			} else if tmp == "lower" {
				style["text-transform"] = "lowercase"
			}
		}
		if len(style) != 0 {
			str = "<span style=\""
			for i, v := range style {
				str += i + ":" + v + ";"
			}
			str += "\">" + meat + "</span>"
		}
	} else if strings.HasPrefix(lwrbb, "color=") {
		col := lwrbb[6:]
		if !strings.HasPrefix(col, "#") {
			_, err := strconv.ParseInt(col, 16, 0)
			if err == nil {
				col = "#" + col
			}
		}
		str = "<span style=\"color:" + col + ";\">" + meat + "</span>"
	} else if strings.HasPrefix(lwrbb, "size=") {
		str = "<span style=\"size:" + lwrbb[5:] + ";\">" + meat + "</span>"
	} else if lwrbb == "smallcaps" {
		str = "<span style=\"font-variant:small-caps;\">" + meat + "</span>"
	} else if lwrbb == "url" || lwrbb == "link" {
		str = "<a href=\"" + meat + "\">" + meat + "</a>"
	} else if strings.HasPrefix(lwrbb, "url") {
		var addr string
		if strings.HasPrefix(lwrbb, "url=") {
			for i := 4; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' {
					addr = lwrbb[4:i]
					break
				}
			}
			if addr == meat {
				addr = lwrbb[4:]
			}
		} else {
			addr = meat
		}
		var title string
		if strings.Contains(lwrbb, "title=") {
			pos := strings.Index(lwrbb, "title=") + 7
			for i := pos; i < len(lwrbb); i++ {
				if lwrbb[i] == lwrbb[pos-1] {
					title = lwrbb[pos:i]
					break
				}
			}
		}
		str = "<a "
		if title != "" {
			if strings.Contains(title, "\"") {
				strings.Replace(title, "\"", "\\\"", -1)
			}
			str += "title=\"" + title + "\" "
		}
		str += "href=\"" + addr + "\">" + meat + "</a>"
	} else if strings.HasPrefix(lwrbb, "link") {
		var addr string
		if strings.HasPrefix(lwrbb, "link=") {
			for i := 5; i < len(lwrbb); i++ {
				if lwrbb[i] == ' ' {
					addr = lwrbb[5:i]
					break
				}
			}
			if addr == meat {
				addr = lwrbb[5:]
			}
		} else {
			addr = meat
		}
		var title string
		if strings.Contains(lwrbb, "title=") {
			pos := strings.Index(lwrbb, "title=") + 7
			for i := pos; i < len(lwrbb); i++ {
				if lwrbb[i] == lwrbb[pos-1] {
					title = lwrbb[pos:i]
					break
				}
			}
		}
		str = "<a "
		if title != "" {
			if strings.Contains(title, "\"") {
				strings.Replace(title, "\"", "\\\"", -1)
			}
			str += "title=\"" + title + "\" "
		}
		str += "href=\"" + addr + "\">" + meat + "</a>"
	} else if lwrbb == "img" {
		str = "<img style='width:20%;' src='" + meat + "'/>"
	} else if strings.HasPrefix(lwrbb, "img") {
		tagness := ""
		style := make(map[string]string)
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
				style["float"] = lf
			} else {
				he := 0
				pos["float"] = strings.Index(lwrbb, "left")
				cnt := strings.Count(lwrbb, "left")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "left")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						style["float"] = lf
						break
					} else {
						tmp = tmp[he+1:]
					}
				}
			}
		} else if strings.Contains(lwrbb, "right") {
			if ((pos["alt"] == 0 || strings.Index(lwrbb, "right") < pos["alt"]) && (pos["title"] == 0 || strings.Index(lwrbb, "right") < pos["title"])) || ((pos["altEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["altEnd"]) && (pos["titleEnd"] == 0 || strings.LastIndex(lwrbb, "right") > pos["titleEnd"])) {
				style["float"] = rt
			} else {
				he := 0
				pos["float"] = strings.Index(lwrbb, "right")
				cnt := strings.Count(lwrbb, "right")
				tmp := lwrbb
				for i := 0; i < cnt; i++ {
					he = he + strings.Index(tmp, "right")
					if (pos["alt"] == 0 || (he < pos["alt"] || he > pos["altEnd"])) && (he == 0 || (he < pos["title"] || he > pos["titleEnd"])) {
						style["float"] = rt
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
		str = "<img" + tagness + " src='" + meat + "'/>"
	} else if lwrbb == "youtube" {
		lwrin := strings.ToLower(meat)
		parsed := ""
		if strings.HasPrefix(lwrin, "http://") || strings.HasPrefix(lwrin, "https://") || strings.HasPrefix(lwrin, "youtu") || strings.HasPrefix(lwrin, "www.") {
			tmp := lwrin[7:]
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
			style["float"] = lf
		}
		if strings.Contains(lwrbb, "right") {
			style["float"] = rt
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
		lwrin := strings.ToLower(meat)
		parsed := ""
		if strings.HasPrefix(lwrin, "http://") || strings.HasPrefix(lwrin, "https://") || strings.HasPrefix(lwrin, "youtu") || strings.HasPrefix(lwrin, "www.") {
			tmp := lwrin[7:]
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
		str = "<iframe"
		if len(style) > 0 {
			str += " style=\""
			for i, v := range style {
				str += i + ":" + v + ";"
			}
			str += "\""
		}
		str += " src='https://www.youtube.com/embed/" + parsed + "' frameborder='0' allowfullscreen></iframe>"
	} else if lwrbb == "ul" {
		meat = strings.Replace(meat, "\n", "", -1)
		spl := strings.Split(meat, "*")
		str = "<ul>"
		for i := 0; i < len(spl); i++ {
			spl[i] = strings.TrimSpace(spl[i])
			if strings.Contains(spl[i], "<ul>") {
				for j := i; j < len(spl); j++ {
					if strings.Contains(spl[j], "</ul>") {
						i = j
					}
				}
			} else if strings.Contains(spl[i], "<ol>") {
				for j := i; j < len(spl); j++ {
					if strings.Contains(spl[j], "</ol>") {
						i = j
					}
				}
			} else if spl[i] != "" {
				str += "<li>" + spl[i] + "</li>"
			}
		}
		str += "</ul>"
		if pWrap {
			str = "</p>" + str + p
		}
	} else if lwrbb == "ol" {
		meat = strings.Replace(meat, "\n", "", -1)
		spl := strings.Split(meat, "*")
		str = "<ol>"
		for i := 0; i < len(spl); i++ {
			spl[i] = strings.TrimSpace(spl[i])
			if strings.Contains(spl[i], "<ul>") {
				for j := i; j < len(spl); j++ {
					if strings.Contains(spl[j], "</ul>") {
						i = j
					}
				}
			} else if strings.Contains(spl[i], "<ol>") {
				for j := i; j < len(spl); j++ {
					if strings.Contains(spl[j], "</ol>") {
						i = j
					}
				}
			} else if spl[i] != "" {
				str += "<li>" + spl[i] + "</li>"
			}
		}
		str += "</ol>"
		if pWrap {
			str = "</p>" + str + p
		}
	} else if lwrbb == "title" {
		meat = strings.Replace(meat, "\n", "", -1)
		str = "<h1>" + meat + "</h1>"
		if pWrap {
			str = "</p>" + str + p
		}
	} else if strings.HasPrefix(lwrbb, "t") {
		meat = strings.Replace(meat, "\n", "", -1)
		par, err := strconv.Atoi(lwrbb[1:])
		if err == strconv.ErrSyntax {
			str = "<h4>" + meat + "</h4>"
		} else {
			if par >= 1 && par <= 6 {
				str = "<h" + strconv.Itoa(par) + ">" + meat + "</h" + strconv.Itoa(par) + ">"
			} else if par < 1 {
				str = "<h1>" + meat + "</h1>"
			} else if par > 6 {
				str = "<h6>" + meat + "</h6>"
			}
		}
		if pWrap {
			str = "</p>" + str + p
		}
	} else if strings.HasPrefix(lwrbb, "align=") {
		str = "<div style=\"text-align:" + lwrbb[6:] + ";\">"
		if pWrap {
			str = "</p>" + str + p + meat + "</p></div>" + p
		} else {
			str = str + meat + "</div>"
		}
	} else {
		return meat
	}
	return str
}
