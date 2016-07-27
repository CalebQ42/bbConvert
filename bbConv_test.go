package bbConvert

import "testing"

var (
	testString = "[b]This is bold[/b]\n Testing paragraph seperating :)\n [CustomTag]This is going to be ignored[/CustomTag] \n [font='Times New Roman' size=20]fonted[/font]"
)

func TestWraping(t *testing.T) {
	AddClass("Classy")
	ImplementDefaults()
	SetWrap(true)
	t.Log(BBtohtml(testString))
}
