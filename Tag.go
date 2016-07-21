package bbConvert

//Tag is a basic bbCode tag (example: [url=http://google.com title="google"])
type Tag struct {
	bbType   string
	isEnd    bool
	params   []string
	values   []string
	fullBB   string
	begIndex int
	endIndex int
}

//FindValue returns the value of a parameter in the tag. If the parameter isn't found
//it returns an empty string. The starting parameter is under "starting". If the
//parameter is lone (has no value) the returned value is equal to the input string.
func (t *Tag) FindValue(param string) string {
	for i, v := range t.params {
		if v == param {
			return t.values[i]
		}
	}
	return ""
}
