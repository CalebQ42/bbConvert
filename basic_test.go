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

func TestYoutube(t *testing.T) {
	testStr := "[youtube]videoid[/youtube][youtube height=200 width=500]http://youtube.com/watch?v=videoid[/youtube][youtube=500x200]http://youtube.com/watch?v=videoid[/youtube][youtube left]youtu.be/videoid[/youtube][youtube right]videoid[/youtube]"
	var expectedarr []string
	//I know it's horribly long, but because the styles get stored in a map then the map is for looped over, the style can be in any order so I have to account for all possibilities.
	expectedarr = append(expectedarr, "<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='width:500;float:none;height:200;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='width:500;float:none;height:200;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>", "<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='width:500;float:none;height:200;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='width:500;float:none;height:200;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>",
		"<p><iframe style='float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='height:200;width:500;float:none;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:none;height:200;width:500;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:left;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe><iframe style='float:right;' src='https://www.youtube.com/embed/videoid' frameborder='0' allowfullscreen></iframe></p>")
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	for _, v := range expectedarr {
		if v == out {
			return
		}
	}
	t.Fatal("Got: " + out)
}

func TestImg(t *testing.T) {
	testStr := "[img=500]image URL[/img] [img right]URL[/img] [img left]URL[/img]"
	var expectedarr []string
	expectedarr = append(expectedarr, "<p><img style='float:none;width:500;' src='image URL'/> <img style='width:20%;float:right;' src='URL'/> <img style='float:left;width:20%;' src='URL'/></p>", "<p><img style='float:none;width:500;' src='image URL'/> <img style='float:right;width:20%;' src='URL'/> <img style='float:left;width:20%;' src='URL'/></p>", "<p><img style='width:500;float:none;' src='image URL'/> <img style='float:right;width:20%;' src='URL'/> <img style='float:left;width:20%;' src='URL'/></p>")
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	for _, v := range expectedarr {
		if v == out {
			return
		}
	}
	t.Fatal("Got:" + out)
}

func TestTitles(t *testing.T) {
	testStr := "[title]TITLE[/title] [t1]ALSO BIG[/t1] [t2]smaller[/t2] [t3]smaller[/t3] [t4]smaller[/t4] [t5]smaller[/t5] [t6]smallest[/t6]"
	expected := "<h1>TITLE</h1><p> <h1>ALSO BIG</h1> <h2>smaller</h2> <h3>smaller</h3> <h4>smaller</h4> <h5>smaller</h5> <h6>smallest</h6></p>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got: " + out + ", Expected: " + expected)
	}
}

func TestAlign(t *testing.T) {
	testStr := "[align=center]centered text[/align]"
	expected := "<div style=\"text-align:center;\"><p>centered text</p></div>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got: " + out + ", Expected: " + expected)
	}
}
