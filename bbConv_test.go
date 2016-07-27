package bbConvert

import "testing"

func TestStandard(t *testing.T) {
	testStr := "[ol][b]Item 1[/b]\n[i]Item 2[/i]\n[font='Times New Roman' size=20 color=red variant=smallcaps]Item 3[/font][/ol]\nNext Paragraph\nAnd the next\n[i]<Orphan tag :)"
	conv := CreateConverter(true, true)
	t.Log(conv.Convert(testStr))
}

func TestNoWraping(t *testing.T) {
	testStr := "[u][i]Hello[s]UnderItalStri[/s][/i][/u][align=center]center[/align][b]<another orphan tag :)"
	conv := CreateConverter(false, true)
	t.Log(conv.Convert(testStr))
	conv.ClearAll()
	t.Log(conv.Convert(testStr))
	conv.ImplementDefaults()
	conv.ClearTag("u")
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

func TestStyleClass(t *testing.T) {
	testStr := "[bold]Simple Bold[/bold] [italics]Simple Italics[/italics]"
	conv := CreateConverter(true, true)
	conv.AddClass("classy")
	conv.AddClass("advclass")
	t.Log(conv.Convert(testStr))
	conv.AddClass(" blahblah")
	conv.SetStyle("height", "50")
	t.Log(conv.Convert(testStr))
	conv.SetStyle("height", "")
	conv.SetStyle("font-family", "Times New Roman")
	t.Log(conv.Convert(testStr))
}

func TestMultipleArguments(t *testing.T) {
	testStr := "[img=500x200 width=200 height=500 title='this is a title' alt='an alternate thing']Image URL[/img]"
	conv := CreateConverter(true, true)
	t.Log(conv.Convert(testStr))
}

func TestMultilevelBullet(t *testing.T) {
	testStr := "[ol]*Lvl1[ol]*lvl2[ol]*lvl3[/ol][/ol][/ol][ul][ol][ul][ol][/ol][/ul][/ol][/ul]"
	conv := CreateConverter(false, true)
	t.Log(conv.Convert(testStr))
	conv.Wrap = true
	t.Log(conv.Convert(testStr))
}

func TestAllTags(t *testing.T) {
	ultimateStr := "[b]Bold[/b][bold]Bold[/bold][i]Italics[/i][italics]Italics[/italics][u]Underline[/u][underline]Underline[/underline][s]Strike[/s][strike]Strike[/strike][font=Verdana]Some Text[/font][font size=20pt]Some Text[/font][font color=red]Some Text[/font][font color=#000000]Some Text[/font][font]Nothing Happening[/font][font color=000000]some text[/font][font variant=upper]Some Text[/font][font variant=lower]Some Text[/font][font variant=smallcaps]Some Text[/font][size=20pt]Some Text[/size][color=red]Some Text[/color][color=#000000]Some Text[/color][smallcaps]Some Text[/smallcaps][url]Link address[/url][url=address]Some Text[/url][url title='Title']Link address[/url][link]Link address[/link][link=address]Some Text[/link][link title='Title']Link address[/link][youtube]Youtube URL or video ID[/youtube][youtube=500]Youtube URL or video ID[/youtube][youtube height=200 width=500]Youtube URL or video ID[/youtube][youtube=500x200]Youtube URL or video ID[/youtube][youtube left]Youtube URL or video ID[/youtube][youtube right]Youtube URL or video ID[/youtube][youtube]http://youtube.com/watch?v=blahblah[/youtube][youtube]blahblah[/youtube][img]Image URL[/img][img=500]Image URL[/img][img=500x200]Image URL[/img][img height=200 width=500]Image URL[/img][img left]Image URL[/img][img right]Image URL[/img][img alt='Alternate text']Image URL[/img][img title='Title']Image URL[/img][image]Image URL[/image][title]Some Text[/title][t1]Some Text[/t1][t2]Some Text[/t2][t3]Some Text[/t3][t4]Some Text[/t4][t5]Some Text[/t5][t6]Some Text[/t6][align=center]Some Text[/align][ul]\n* Item 1\nItem 2\n[/ul][ol]\n* Item 1\nItem 2\n[/ol][bullet] * Item 1 * Item 2[/bullet][number] * Item 1 * Item 2[/number][ul]* Item 1 * Item 2[/ul][ol]* Item 1 * Item 2[/ol]"
	conv := CreateConverter(true, true)
	t.Log(conv.Convert(ultimateStr))
}
