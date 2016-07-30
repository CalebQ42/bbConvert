package bbConvert

import (
	"strconv"
	"strings"
)

//HTMLConverter is an easy way to convert bbCode to HTML. It automatically wraps the output
//in paragraph tags and properly converts newlines (\n) to paragraphs.
type HTMLConverter struct {
	conv  Converter
	p     string
	class []string
	style map[string]string
}

//Convert converts the input with bbCode to output with HTML
func (h *HTMLConverter) Convert(in string) string {
	out := h.conv.Convert(in)
	out = strings.Replace(out, "\n", "</p>"+h.StartingParagraphTag(), -1)
	out = h.StartingParagraphTag() + out + "</p>"
	return out
}

//StartingParagraphTag returns the starting paragraph tags used when wraping the output in paragraph
//tags with the proper style and class(es)
func (h *HTMLConverter) StartingParagraphTag() string {
	p := "<p"
	if h.style != nil && len(h.style) != 0 {
		p += " style=\""
		for i, v := range h.style {
			p += i + ":"
			if strings.Contains(v, " ") {
				v = "'" + v + "'"
			}
			p += v + ";"
		}
		p += "\""
	}
	if len(h.class) != 0 {
		p += " class=\"" + strings.Join(h.class, " ") + "\""
	}
	p += ">"
	return p
}

//AddClass adds a class to the paragraph tags used to wrap the output. Multiple
//classes can be added at once if they are seperated by spaces.
func (h *HTMLConverter) AddClass(class string) {
	if strings.Contains(class, " ") {
		spl := strings.Split(class, " ")
		h.class = append(h.class, spl...)
	} else {
		h.class = append(h.class, class)
	}
}

//SetStyle sets a give style to the paragraph tags used to wrap the output.
func (h *HTMLConverter) SetStyle(css, value string) {
	if h.style == nil {
		h.style = make(map[string]string)
	}
	if value == "" {
		delete(h.style, css)
	} else {
		h.style[css] = value
	}
}

//ImplementDefaults adds the default supported bbCode to HTML conversions.
func (h *HTMLConverter) ImplementDefaults() {
	if h.conv.funcs == nil {
		h.conv.funcs = make(map[string]func(Tag, string) string)
	}
	tmp := func(_ Tag, meat string) string {
		return "<b>" + meat + "</b>"
	}
	h.conv.funcs["b"] = tmp
	h.conv.funcs["bold"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<i>" + meat + "</i>"
	}
	h.conv.funcs["i"] = tmp
	h.conv.funcs["italics"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<s>" + meat + "</s>"
	}
	h.conv.funcs["s"] = tmp
	h.conv.funcs["strike"] = tmp
	tmp = func(_ Tag, meat string) string {
		return "<u>" + meat + "</u>"
	}
	h.conv.funcs["u"] = tmp
	h.conv.funcs["underline"] = tmp
	h.conv.funcs["font"] = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["color"] = fnt.Value("color")
		style["font-family"] = fnt.Value("starting")
		style["font-size"] = fnt.Value("size")
		tmp := fnt.Value("variant")
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
		for i, v := range style {
			if v == "" {
				delete(style, i)
			}
		}
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
	h.conv.funcs["color"] = func(fnt Tag, meat string) string {
		clr := fnt.Value("starting")
		if clr != "" {
			if !strings.HasPrefix(clr, "#") {
				_, err := strconv.ParseInt(clr, 16, 0)
				if err == nil {
					clr = "#" + clr
				}
			}
			return "<span style='color:" + clr + ";'>" + meat + "</span>"
		}
		return meat
	}
	h.conv.funcs["size"] = func(fnt Tag, meat string) string {
		return "<span style='font-color:" + fnt.Value("starting") + ";'>" + meat + "</span>"
	}
	h.conv.funcs["smallcaps"] = func(_ Tag, meat string) string {
		return "<span style=\"font-variant:small-caps;\">" + meat + "</span>"
	}
	tmp = func(fnt Tag, meat string) string {
		out := "<a"
		if fnt.Value("title") != "" {
			out += " title='" + strings.Replace(fnt.Value("title"), "'", "\\'", -1) + "'"
		}
		if fnt.Value("starting") == "" {
			out += " href=\"" + meat + "\">" + meat + "</a>"
		} else {
			out += " href=\"" + fnt.Value("starting") + "\">" + meat + "</a>"
		}
		return out
	}
	h.conv.funcs["url"] = tmp
	h.conv.funcs["link"] = tmp
	tmp = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.Value("right") != "" {
			style["float"] = fnt.Value("right")
		} else if fnt.Value("left") != "" {
			style["float"] = fnt.Value("left")
		}
		alt := fnt.Value("alt")
		title := fnt.Value("title")
		style["width"] = fnt.Value("width")
		style["height"] = fnt.Value("height")
		srt := fnt.Value("starting")
		if srt != "" {
			if ind := strings.Index(srt, "x"); ind != -1 {
				style["width"] = srt[:ind]
				style["height"] = srt[ind+1:]
			} else {
				style["width"] = srt
			}
		}
		out = "<img "
		for i, v := range style {
			if v == "" {
				delete(style, i)
			}
		}
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
	h.conv.funcs["img"] = tmp
	h.conv.funcs["image"] = tmp
	h.conv.funcs["youtube"] = func(fnt Tag, meat string) string {
		var out string
		style := make(map[string]string)
		style["float"] = "none"
		if fnt.Value("right") != "" {
			style["float"] = fnt.Value("right")
		} else if fnt.Value("left") != "" {
			style["float"] = fnt.Value("left")
		}
		if fnt.Value("height") != "" {
			style["height"] = fnt.Value("height")
		}
		if fnt.Value("width") != "" {
			style["width"] = fnt.Value("width")
		}
		if fnt.Value("starting") != "" {
			if strings.Contains(fnt.Value("starting"), "x") {
				spl := strings.Split(fnt.Value("starting"), "x")
				style["height"] = spl[1]
				style["width"] = spl[0]
			} else {
				style["width"] = fnt.Value("starting")
			}
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
		out = "<iframe"
		if len(style) != 0 {
			out += " style='"
			for i, v := range style {
				out += i + ":" + v + ";"
			}
			out += "'"
		}
		out += " src='https://www.youtube.com/embed/" + parsed + "' frameborder='0' allowfullscreen></iframe>"
		return out
	}
	h.conv.funcs["title"] = func(fnt Tag, meat string) string {
		var out string
		meat = strings.Replace(meat, "\n", "", -1)
		out = "<h1>" + meat + "</h1>"
		out = "</p>" + out + h.StartingParagraphTag()
		return out
	}
	h.conv.funcs["align"] = func(fnt Tag, meat string) string {
		var out string
		out = "<div style=\"text-align:" + fnt.Value("starting") + ";\">"
		out = "</p>" + out + h.StartingParagraphTag() + meat + "</p></div>" + h.StartingParagraphTag()
		return out
	}
	h.conv.funcs["t1"] = func(fnt Tag, meat string) string {
		return "<h1>" + meat + "</h1>"
	}
	h.conv.funcs["t2"] = func(fnt Tag, meat string) string {
		return "<h2>" + meat + "</h2>"
	}
	h.conv.funcs["t3"] = func(fnt Tag, meat string) string {
		return "<h3>" + meat + "</h3>"
	}
	h.conv.funcs["t4"] = func(fnt Tag, meat string) string {
		return "<h4>" + meat + "</h4>"
	}
	h.conv.funcs["t5"] = func(fnt Tag, meat string) string {
		return "<h5>" + meat + "</h5>"
	}
	h.conv.funcs["t6"] = func(fnt Tag, meat string) string {
		return "<h6>" + meat + "</h6>"
	}
	//re-do lists
}

//Converter returns the Converter that's used in the HTMLConverter
//so you can add custom functions and other items that you can access with the Converter
func (h *HTMLConverter) Converter() *Converter {
	return &h.conv
}
