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
		pos := make(map[string][]int)
		stuff := make(map[string]int)
		pos["alt"] = indexAll(lwrbb, "alt=")
		pos["title"] = indexAll(lwrbb, "title=")
		pos["width"] = indexAll(lwrbb, "width=")
		pos["height"] = indexAll(lwrbb, "height=")
		pos["left"] = indexAll(lwrbb, "left")
		pos["right"] = indexAll(lwrbb, "right")
		if len(pos["alt"]) != 0 || len(pos["title"]) != 0 {
			if len(pos["title"]) != 0 && (len(pos["alt"]) == 0 || pos["alt"][0] > pos["title"][0]) {
				stuff["title"] = 0
				posi := pos["title"][0] + 7
				qt := lwrbb[posi-1]
				for i := posi; i < len(bb); i++ {
					if lwrbb[i] == qt {
						other["title"] = bb[posi:i]
						break
					}
				}
				if other["title"] == "" {
					other["title"] = bb[posi:]
				}
				leng := len(other["title"])
				for i, v := range pos["alt"] {
					if v > pos["title"][0]+leng+7 {
						stuff["title"] = i
						posi = v + 5
						qt = lwrbb[posi-1]
						for j := posi; j < len(bb); j++ {
							if lwrbb[j] == qt {
								other["alt"] = bb[posi:j]
								break
							}
						}
						if other["alt"] == "" {
							other["alt"] = bb[posi:]
						}
						break
					}
				}
			} else if len(pos["alt"]) != 0 && (len(pos["title"]) == 0 || pos["alt"][0] < pos["title"][0]) {
				stuff["alt"] = 0
				posi := pos["alt"][0] + 5
				qt := lwrbb[posi-1]
				for i := posi; i < len(bb); i++ {
					if lwrbb[i] == qt {
						other["alt"] = bb[posi:i]
						break
					}
				}
				if other["alt"] == "" {
					other["alt"] = bb[posi:]
				}
				leng := len(other["alt"])
				for i, v := range pos["title"] {
					if v > pos["alt"][0]+leng+5 {
						stuff["title"] = i
						posi = v + 7
						qt = lwrbb[posi-1]
						for j := posi; j < len(bb); j++ {
							if lwrbb[j] == qt {
								other["title"] = bb[posi:j]
								break
							}
						}
						if other["title"] == "" {
							other["title"] = bb[posi:]
						}
						break
					}
				}
			}
		}
		if len(pos["width"]) != 0 {
			for _, v := range pos["width"] {
				if other["title"] == "" || (v < pos["title"][stuff["title"]] || v > pos["title"][stuff["title"]]+len(other["title"])+7) {
					if other["alt"] == "" || (v < pos["alt"][stuff["alt"]] || v > pos["alt"][stuff["alt"]]+len(other["title"])+5) {
						posi := v + 6
						for j := posi; j < len(bb); j++ {
							if bb[j] == ' ' {
								style["width"] = bb[posi:j]
								break
							} else if j == len(bb)-1 {
								style["width"] = bb[posi:]
							}
						}
						break
					}
				}
			}
		}
		if len(pos["height"]) != 0 {
			for _, v := range pos["height"] {
				if other["title"] == "" || (v < pos["title"][stuff["title"]] || v > pos["title"][stuff["title"]]+len(other["title"])+7) {
					if other["alt"] == "" || (v < pos["alt"][stuff["alt"]] || v > pos["alt"][stuff["alt"]]+len(other["title"])+5) {
						posi := v + 7
						for j := posi; j < len(bb); j++ {
							if bb[j] == ' ' {
								style["height"] = bb[posi:j]
								break
							} else if j == len(bb)-1 {
								style["height"] = bb[posi:]
							}
						}
						break
					}
				}
			}
		}
		if len(pos["left"]) != 0 {
			for _, v := range pos["left"] {
				if other["title"] == "" || (v < pos["title"][stuff["title"]] || v > pos["title"][stuff["title"]]+len(other["title"])+7) {
					if other["alt"] == "" || (v < pos["alt"][stuff["alt"]] || v > pos["alt"][stuff["alt"]]+len(other["title"])+5) {
						style["float"] = "left"
						break
					}
				}
			}
		}
		if len(pos["right"]) != 0 {
			for _, v := range pos["right"] {
				if other["title"] == "" || (v < pos["title"][stuff["title"]] || v > pos["title"][stuff["title"]]+len(other["title"])+7) {
					if other["alt"] == "" || (v < pos["alt"][stuff["alt"]] || v > pos["alt"][stuff["alt"]]+len(other["title"])+5) {
						style["float"] = "right"
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
		str = "<img" + tagness + " src='" + strings.TrimSpace(meat) + "'/>"
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
