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

func TestAllTags(t *testing.T) {
	ultimateStr := "[b]Bold[/b][bold]Bold[/bold][i]Italics[/i][italics]Italics[/italics][u]Underline[/u][underline]Underline[/underline][s]Strike[/s][strike]Strike[/strike][font=Verdana]Some Text[/font][font size=20pt]Some Text[/font][font color=red]Some Text[/font][font color=#000000]Some Text[/font][font variant=upper]Some Text[/font][font variant=lower]Some Text[/font][font variant=smallcaps]Some Text[/font][size=20pt]Some Text[/size][color=red]Some Text[/color][color=#000000]Some Text[/color][smallcaps]Some Text[/smallcaps][url]Link address[/url][url=address]Some Text[/url][url title='Title']Link address[/url][link]Link address[/link][link=address]Some Text[/link][link title='Title']Link address[/link][youtube]Youtube URL or video ID[/youtube][youtube height=200 width=500]Youtube URL or video ID[/youtube][youtube=500x200]Youtube URL or video ID[/youtube][youtube left]Youtube URL or video ID[/youtube][youtube right]Youtube URL or video ID[/youtube][img]Image URL[/img][img=500x200]Image URL[/img][img height=200 width=500]Image URL[/img][img left]Image URL[/img][img right]Image URL[/img][img alt='Alternate text']Image URL[/img][img title='Title']Image URL[/img][image]Image URL[/image][title]Some Text[/title][t1]Some Text[/t1][t2]Some Text[/t2][t3]Some Text[/t3][t4]Some Text[/t4][t5]Some Text[/t5][t6]Some Text[/t6][align=center]Some Text[/align][ul]\n* Item 1\nItem 2\n[/ul][ol]\n* Item 1\nItem 2\n[/ol][bullet] * Item 1 * Item 2[/bullet][number] * Item 1 * Item 2[/number][ul]* Item 1 * Item 2[/ul][ol]* Item 1 * Item 2[/ol]"
	conv := CreateConverter(true, true)
	t.Log(conv.Convert(ultimateStr))
	//Should output:
	//
	//<p><b>Bold</b><b>Bold</b><i>Italics<i><i>Italics<i><u>Underline</u><u>Underline</u><s>Strike</s><s>Strike</s><span style="font-family:Verdana;">Some Text</span><span style="font-size:20pt;">Some Text</span><span style="color:red;">Some Text</span><span style="color:#000000;">Some Text</span><span style="text-transform:uppercase;">Some Text</span><span style="text-transform:lowercase;">Some Text</span><span style="font-variant:small-caps;">Some Text</span><span style='font-color:20pt;'>Some Text</span><span style='color:red;'>Some Text</span><span style='color:#000000;'>Some Text</span><span style="font-variant:small-caps;">Some Text</span><a href="Link address">Link address</a><a href="address">Some Text</a><a href="Link address">Link address</a><a href="Link address">Link address</a><a href="address">Some Text</a><a href="Link address">Link address</a><iframstyle='float:none;'  src='https://www.youtube.com/embed/ URL or video ID' frameborder='0' allowfullscreen></iframe><iframstyle='float:none;'  src='https://www.youtube.com/embed/ URL or video ID' frameborder='0' allowfullscreen></iframe><iframstyle='float:none;'  src='https://www.youtube.com/embed/ URL or video ID' frameborder='0' allowfullscreen></iframe><iframstyle='float:left;'  src='https://www.youtube.com/embed/ URL or video ID' frameborder='0' allowfullscreen></iframe><iframstyle='float:right;'  src='https://www.youtube.com/embed/ URL or video ID' frameborder='0' allowfullscreen></iframe><img style='float:none;width:20%;' src='Image URL'/><img style='float:none;width:500;height:200;' src='Image URL'/><img style='float:none;width:500;height:200;' src='Image URL'/><img style='float:left;width:20%;' src='Image URL'/><img style='width:20%;float:right;' src='Image URL'/><img style='float:none;width:20%;' alt='Alternate text' src='Image URL'/><img style='float:none;width:20%;' title='Title' src='Image URL'/><img style='float:none;width:20%;' src='Image URL'/></p><h1>Some Text</h1><p><h1>Some Text</h1><h2>Some Text</h2><h3>Some Text</h3><h4>Some Text</h4><h5>Some Text</h5><h6>Some Text</h6></p><div style="text-align:center;"><p>Some Text</p></div><ul><li>Item 1</li><li>Item 2</li></ul><ol><li>Item 1</li><li>Item 2</li></ol><ul><li>Item 1</li><li>Item 2</li></ul><ol><li>Item 1</li><li>Item 2</li></ol><ul><li>Item 1</li><li>Item 2</li></ul><ol><li>Item 1</li><li>Item 2</li></ol>

}
