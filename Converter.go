//Package bbConvert provides an easy way to process and convert bbCode. HTMLConverter
//is an easier way to convert to HTML.
package bbConvert

import "strings"

//Converter provides an easy way to convert bbCode.
type Converter struct {
	funcs map[string]func(Tag, string) string
}

//AddCustom adds a conversion function to the Converter. bbType is self explanitory
//(it is case insensitive) and the function is type func(Tag,string) string: The first
//input is the front Tag and the second input in the text found in between the two matched tags
func (c *Converter) AddCustom(bbType string, f func(Tag, string) string) {
	c.funcs[strings.ToLower(bbType)] = f
}

func (c Converter) conv(tag Tag, meat string) string {
	if _, ok := c.funcs[tag.typ]; ok {
		return c.funcs[tag.typ](tag, meat)
	}
	return ""
}
