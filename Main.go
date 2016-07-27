/*Package bbConvert is used to convert bbCode to HTML. Has support for adding support for custom bbCode tags.*/
package bbConvert

import "strings"

//Converter processes and converts bbCode tags in a string
type Converter struct {
	Wrap     bool
	pStyle   map[string]string
	pClass   string
	p        string
	tagFuncs map[string]func(Tag, string) string
}

//CreateConverter creates a Convert. The first bool sets if the Converter's output should be wrapped in paragraph tags. The second bool sets if the bbCode to HTML defaults are used.
func CreateConverter(wrap bool, defaults bool) (out Converter) {
	out.Wrap = wrap
	if defaults {
		if out.tagFuncs == nil {
			out.tagFuncs = make(map[string]func(Tag, string) string)
		}
		out.ImplementDefaults()
	}
	return
}

//AddCustomTag allows you to add support for a custom bbCode tag.
func (c *Converter) AddCustomTag(tag string, f func(Tag, string) string) {
	if c.tagFuncs == nil {
		c.tagFuncs = make(map[string]func(Tag, string) string)
	}
	c.tagFuncs[tag] = f
}

//AddClass adds a class to the paragraph tags wrapped around the output if Wrap is set to true. Can be added singularly or with space separation between multiple classes.
func (c *Converter) AddClass(cl string) {
	cl = strings.TrimSpace(cl)
	if c.pClass == "" {
		c.pClass += cl
	} else {
		if strings.HasSuffix(c.pClass, " ") {
			c.pClass += cl
		} else {
			c.pClass += " " + cl
		}
	}
	c.updatep()
}

//StartingParagraphTag returns the properly processed with the set style and class. If Wrap = false then returns an empty string.
func (c *Converter) StartingParagraphTag() string {
	if c.p == "" && c.Wrap == true {
		c.updatep()
	}
	return c.p
}

//SetStyle adds a style to the paragraph tags wrapped around the output if Wrap is set true. If the style value is set to an empty string then it removes the style.
func (c *Converter) SetStyle(css, value string) {
	if c.pStyle == nil {
		c.pStyle = make(map[string]string)
	}
	if value != "" {
		value = strings.TrimSpace(value)
		value = strings.Trim(value, "'")
		value = strings.Trim(value, "\"")
		c.pStyle[css] = value
	} else {
		if c.pStyle != nil {
			_, ok := c.pStyle[css]
			if ok {
				delete(c.pStyle, css)
				if len(c.pStyle) == 0 {
					c.pStyle = nil
				}
			}
		}
	}
	c.updatep()
}

//Convert actually converts the input string and returns the converted string
func (c Converter) Convert(input string) string {
	c.updatep()
	input = strings.Trim(input, "\n")
	out := c.conv(input)
	out = "\n" + out + "\n"
	if c.Wrap {
		out = strings.Replace(out, "\n", "</p>"+c.p, -1)
		out = strings.TrimPrefix(out, "</p>")
		out = strings.TrimSuffix(out, c.p)
		out = strings.Replace(out, c.p+"</p>", "", -1)
	} else {
		out = strings.Replace(out, "\n", " ", -1)
	}
	return out
}

func (c Converter) conv(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			for j := i; j < len(input); j++ {
				if input[j] == ']' {
					tmpTag := processTag(input[i : j+1])
					if !tmpTag.isEnd {
						tmpTag.begIndex = i
						tmpTag.endIndex = j
						ndTag := findEndTag(tmpTag, input)
						if ndTag.bbType != "" {
							out := c.convTag(tmpTag, ndTag, c.conv(input[tmpTag.endIndex+1:ndTag.begIndex]))
							input = input[:i] + out + input[ndTag.endIndex+1:]
						}
					}
					break
				}
			}
		}
	}
	return input
}

func (c *Converter) updatep() {
	if c.Wrap {
		c.p = "<p"
		if c.pStyle != nil && len(c.pStyle) != 0 {
			c.p += " style=\""
			for i, v := range c.pStyle {
				if strings.Contains(v, " ") {
					v = "'" + v + "'"
				}
				c.p += i + ":" + v + ";"
			}
			c.p += "\""
		}
		if c.pClass != "" {
			c.p += " class=\"" + c.pClass + "\""
		}
		c.p += ">"
	} else {
		c.p = ""
	}
}

func (c Converter) convTag(beg, end Tag, meat string) string {
	if _, ok := c.tagFuncs[beg.bbType]; ok {
		return c.tagFuncs[beg.bbType](beg, meat)
	}
	return beg.fullBB + meat + end.fullBB
}
