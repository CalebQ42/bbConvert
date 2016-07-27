package bbConvert

import "testing"

func TestStandard(t *testing.T) {
	testStr := "[ol][b]Item 1[/b]\n[i]Item 2[/i]\n[font='Times New Roman' size=20 color=red variant=smallcaps]Item 3[/font][/ol]\nNext Paragraph\nAnd the next\n[i]<Orphan tag :)"
	conv := CreateConverter(true, true)
	t.Log(conv.Convert(testStr))
}

func TestNoWraping(t *testing.T) {
	testStr := "[u][i]Hello[s]UnderItalStri[/s][/i][/u][b]<another orphan tag :)"
	conv := CreateConverter(false, true)
	t.Log(conv.Convert(testStr))
}

func TestMarkdown(t *testing.T) {
	testStr := "[bold]Simple Bold[/bold] [italics]Simple Italics[/italics]"
	conv := CreateConverter(false, false)
	conv.AddCustomTag("bold", func(fnt Tag, meat string) string {
		return "**" + meat + "**"
	})
	conv.AddCustomTag("italics", func(fnt Tag, meat string) string {
		return "*" + meat + "*"
	})
	t.Log(conv.Convert(testStr))
}
