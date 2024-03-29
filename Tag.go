package bbConvert

import "strings"

// Tag is a bbCode tag that is properly processed.
type Tag struct {
	params map[string]string
	typ    string
	index  [2]int
	end    bool
}

// Value returns the value of a parameter of the bbCode tag. Starting parameters
// is under "starting". Parameter names are always in lowercase
func (t *Tag) Value(param string) string {
	return t.params[param]
}

func (t *Tag) setValue(param, value string) {
	if t.params == nil {
		t.params = make(map[string]string)
	}
	t.params[strings.TrimSpace(strings.ToLower(param))] = strings.TrimSpace(value)
}

func (t *Tag) process(bbtag string) {
	if strings.HasPrefix(bbtag, "/") {
		t.end = true
		t.typ = strings.ToLower(strings.TrimPrefix(bbtag[:len(bbtag)-1], "/"))
		return
	}
	for i, v := range bbtag {
		if v == '=' || v == ' ' || v == ']' {
			t.typ = strings.ToLower(bbtag[:i])
			switch v {
			case '=':
				if qt := bbtag[i+1]; qt == '\'' || qt == '"' {
					for j := i + 2; j < len(bbtag); j++ {
						if bbtag[j] == qt {
							t.setValue("starting", bbtag[i+2:j])
							bbtag = bbtag[j+1:]
							break
						} else if bbtag[j] == ']' {
							t.setValue("starting", bbtag[i+2:j])
							return
						}
					}
				} else {
					for j := i + 1; j < len(bbtag); j++ {
						if bbtag[j] == ']' {
							t.setValue("starting", bbtag[i+1:j])
							return
						} else if bbtag[j] == ' ' {
							t.setValue("starting", bbtag[i+1:j])
							bbtag = bbtag[j+1:]
							break
						}
					}
				}
			case ']':
				return
			case ' ':
				bbtag = bbtag[i:]
			}
			break
		}
	}
	t.processFurther(bbtag)
}

func (t *Tag) processFurther(further string) {
	further = strings.TrimSpace(further)
	for i := 0; i < len(further); i++ {
		switch further[i] {
		case ' ':
			t.setValue(strings.ToLower(further[:i]), further[:i])
			further = strings.TrimSpace(further[i:])
			i = -1
		case '=':
			if qt := further[i+1]; qt == '\'' || qt == '"' {
			outloopqt:
				for j := i + 2; j < len(further); j++ {
					switch further[j] {
					case ']':
						t.setValue(strings.ToLower(further[:i]), further[i+2:j])
						return
					case qt:
						t.setValue(strings.ToLower(further[:i]), further[i+2:j])
						further = strings.TrimSpace(further[j+1:])
						i = -1
						break outloopqt
					}
				}
			} else {
			outloop:
				for j := i + 1; j < len(further); j++ {
					switch further[j] {
					case ']':
						t.setValue(strings.ToLower(further[:i]), further[i+1:j])
						return
					case ' ':
						t.setValue(strings.ToLower(further[:i]), further[i+1:j])
						further = strings.TrimSpace(further[j:])
						i = -1
						break outloop
					}
				}
			}
		case ']':
			if i != 0 {
				t.setValue(strings.ToLower(further[:i]), further[:i])
				return
			}
			return
		}
	}
}
