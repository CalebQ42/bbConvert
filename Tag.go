package bbConvert

type Tag struct {
	bbType   string
	isEnd    bool
	params   []string
	values   []string
	fullBB   string
	begIndex int
	endIndex int
}

func (t *Tag) FindValue(param string) string {
	for i, v := range t.params {
		if v == param {
			return t.values[i]
		}
	}
	return ""
}
