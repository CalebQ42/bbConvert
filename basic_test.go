package bbConvert

import "testing"

func TestBasicTags(t *testing.T) {
	testStr := "[b]Bold[/b] [bold]Also Bold[/bold] [i]Italics[/i] [italics]Italics[/italics] [u]Underline[/u] [underline]Underline[/underline] [s]Strike[/s] [strike]Strike[/strike]"
	expected := "<p><b>Bold</b> <b>Also Bold</b> <i>Italics<i> <i>Italics<i> <u>Underline</u> <u>Underline</u> <s>Strike</s> <s>Strike</s></p>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got:" + out + ", Expected:" + expected)
	}
}

func TestFont(t *testing.T) {
	testStr := "[font=Verdana]Texty[/font] [font='Times New Romans']Texty More[/font][font size=20] Bigger is better[/font] [font color=red]redness[/font] [font color=545454]some random color[/font] [font color=#545454]some random color[/font] [font variant=upper]thiss will be in caps[/font] [font variant=lower]AND THIS WILL BE IN LOWERCASE[/font] [font variant=smallcaps]and, of course, small caps[/font]"
	expected := "<p><span style=\"font-family:Verdana;\">Texty</span> <span style=\"font-family:'Times New Romans';\">Texty More</span><span style=\"font-size:20;\"> Bigger is better</span> <span style=\"color:red;\">redness</span> <span style=\"color:#545454;\">some random color</span> <span style=\"color:#545454;\">some random color</span> <span style=\"text-transform:uppercase;\">thiss will be in caps</span> <span style=\"text-transform:lowercase;\">AND THIS WILL BE IN LOWERCASE</span> <span style=\"font-variant:small-caps;\">and, of course, small caps</span></p>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got:" + out + ", Expected:" + expected)
	}
}

func TestSerperateStyle(t *testing.T) {
	testStr := "[size=20]Bigger is Better[/size] [color=red]Reder[/color] [color=545454]RANDOM COLOR[/color] [smallcaps][color=#545454]small caps in some random color[/color][/smallcaps]"
	expected := "<p><span style='font-color:20;'>Bigger is Better</span> <span style='color:red;'>Reder</span> <span style='color:#545454;'>RANDOM COLOR</span> <span style=\"font-variant:small-caps;\"><span style='color:#545454;'>small caps in some random color</span></span></p>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got:" + out + ", Expected:" + expected)
	}
}

func TestLinkURL(t *testing.T) {
	testStr := "[url]google.com[/url] [link=google.com]A link to google[/link] [url=google.com title='Blah Blah']A link to google[/url]"
	expected := "<p><a href=\"google.com\">google.com</a> <a href=\"google.com\">A link to google</a> <a title='Blah Blah' href=\"google.com\">A link to google</a></p>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got:" + out + ", Expected:" + expected)
	}
}
