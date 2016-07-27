package bbConvert

import "testing"

func TestWraping(t *testing.T) {
	AddClass("Classy")
	ImplementDefaults()
	SetWrap(true)
	AddCustomTag("customtag", func(fnt Tag, meat string) string {
		return "This is custom!"
	})
	t.Log(BBtohtml("[b]This is bold[/b]\n Testing paragraph seperating :)\n [CustomTag]This is going to be ignored[/CustomTag] \n [font='Times New Roman' size=20]fonted[/font]"))
}

func TestAlternate(t *testing.T) {
	SetWrap(false)
	AddCustomTag("b", func(fnt Tag, meat string) string {
		return "**" + meat + "**"
	})
	AddCustomTag("i", func(fnt Tag, meat string) string {
		return "*" + meat + "*"
	})
	t.Log(BBtohtml("[b]This should be in markdown style bold[/b] [i]And now markdown italics[/i]"))
}
