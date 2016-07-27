package bbConvert

import "testing"

var (
	testString = "[b]This is bold[/b]\n Testing paragraph seperating :)\n [CustomTag]This is going to be ignored[/CustomTag]"
)

func TestWraping(t *testing.T) {
	AddClass("Classy")
	ImplementDefaults()
	AddCustomTag("customtag", func(fnt Tag, meat string) string {
		return "I'm not caring about the meat right now"
	})
	SetWrap(true)
	t.Log(BBtohtml(testString))
}
