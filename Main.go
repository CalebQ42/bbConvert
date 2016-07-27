/*Package bbConvert is used to convert bbCode to HTML. Has support for adding support for custom bbCode tags.*/
package bbConvert

import "strings"

var (
	pWrap = false
	style map[string]string
	class string
	p     string
)

//BBtohtml simply takes in a string with BBCode and exports the string converted to html
func BBtohtml(input string) string {
	updatep()
	input = strings.Trim(input, "\n")
	out := bbconv(input)
	out = "\n" + out + "\n"
	if pWrap {
		out = strings.Replace(out, "\n", "</p>"+p, -1)
		out = strings.TrimPrefix(out, "</p>")
		out = strings.TrimSuffix(out, p)
		out = strings.Replace(out, p+"</p>", "", -1)
	} else {
		out = strings.Replace(out, "\n", " ", -1)
	}
	return out
}

//SetWrap determines if the output should have pargraph tags at every newline
//If it's set to false(default) then newlines will be converted to spaces
func SetWrap(tf bool) {
	pWrap = tf
	updatep()
}

//SetStyle allows you to add a particular css style to the wraped paragraph tags
//If you want to undo a style just set it's value to an empty string
//Don't worry about the value having white spaces, it will automagically add
//single quotes if there's whitespace
func SetStyle(css, value string) {
	if value != "" {
		value = strings.TrimSpace(value)
		value = strings.Trim(value, "'")
		value = strings.Trim(value, "\"")
		if style == nil {
			style = make(map[string]string)
		}
		style[css] = value
	} else {
		if style != nil {
			_, ok := style[css]
			if ok {
				delete(style, css)
				if len(style) == 0 {
					style = nil
				}
			}
		}
	}
	updatep()
}

//AddClass allows you to add a particular class to the wraped paragraph tags
func AddClass(cl string) {
	if class == "" {
		class += cl
	} else {
		if strings.HasSuffix(class, " ") {
			class += cl
		} else {
			class += " " + cl
		}
	}
	updatep()
}

func updatep() {
	if pWrap {
		p = "<p"
		if style != nil {
			p += " style=\""
			for i, v := range style {
				if strings.Contains(v, " ") {
					v = "'" + v + "'"
				}
				p += i + ":" + v + ";"
			}
			p += "\""
		}
		if class != "" {
			p += " class=\"" + class + "\""
		}
		p += ">"
	} else {
		p = ""
	}
}
