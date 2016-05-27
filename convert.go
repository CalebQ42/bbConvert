package bbConvert

import "strings"

type tag struct {
	bbType   string
	isEnd    bool
	params   []string
	values   []string
	fullBB   string
	begIndex int
	endIndex int
}

func (t *tag) findValue(param string) string {
	for i, v := range t.params {
		if v == param {
			return t.values[i]
		}
	}
	return ""
}

func bbconv(input string) string {
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
							out := toHTML(tmpTag, ndTag, bbconv(input[tmpTag.endIndex+1:ndTag.begIndex]))
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

func findEndTag(fnt tag, str string) tag {
	var count int
	for i := fnt.endIndex + 1; i < len(str); i++ {
		if str[i] == '[' {
			for j := i; j < len(str); j++ {
				if str[j] == ']' {
					tmpTag := processTag(str[i : j+1])
					if tmpTag.bbType == fnt.bbType {
						if tmpTag.isEnd {
							count--
							if count == -1 {
								tmpTag.begIndex = i
								tmpTag.endIndex = j
								return tmpTag
							}
						} else {
							count++
							break
						}
					}
				}
			}
		}
	}
	return tag{}
}

func processTag(str string) (out tag) {
	out.fullBB = str
	str = str[1:]
	if strings.HasPrefix(str, "/") {
		out.isEnd = true
		out.bbType = strings.ToLower(str[1 : len(str)-1])
		return
	}
	for i, v := range str {
		if v == ']' || v == ' ' || v == '=' {
			out.bbType = strings.ToLower(str[:i])
			if v == ']' {
				return
			} else if v == '=' {
				if str[i+1] == '\'' || str[i+1] == '"' {
					qt := str[i+1]
					for j := i + 2; j < len(str); j++ {
						if str[j] == ']' || str[j] == qt {
							out.params = append(out.params, "starting")
							out.values = append(out.values, str[i+2:j])
							if str[j] == ']' {
								return
							}
							str = str[j+1:]
							break
						}
					}
					break
				} else {
					for j := i + 1; j < len(str); j++ {
						if str[j] == ']' || str[j] == ' ' {
							out.params = append(out.params, "starting")
							out.values = append(out.values, str[i+1:j])
							if str[j] == ']' {
								return
							}
							str = str[j+1:]
							break
						}
					}
					break
				}
			}
		}
	}
	str = strings.TrimSpace(str)
	var prev int
	for i := 0; i < len(str); i++ {
		v := str[i]
		if v == '=' {
			out.params = append(out.params, strings.ToLower(str[prev:i]))
			if str[i+1] == '\'' || str[i+1] == '"' {
				qt := str[i+1]
				for j := i + 2; j < len(str); j++ {
					if str[j] == ']' || str[j] == qt {
						out.values = append(out.values, str[i+2:j])
						if str[j] == ']' || str[j+1] == ']' {
							return
						}
						i = j + 2
						prev = j + 2
						break
					}
				}
			} else {
				for j := i + 1; j < len(str); j++ {
					if str[j] == ']' || str[j] == ' ' {
						out.values = append(out.values, str[i+1:j])
						if str[j] == ']' {
							return
						}
						i = j + 1
						prev = j + 1
						break
					}
				}
			}
		} else if v == ' ' || v == ']' {
			out.params = append(out.params, strings.ToLower(str[prev:i]))
			out.values = append(out.values, str[prev:i])
			if v == ']' {
				return
			}
			prev = i + 1
		}
	}
	return
}
