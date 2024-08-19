package bbConvert

import (
	"testing"
)

const (
	bbTestString = `[b]Some Text[/b] //bolded text
[bold]Some Text[/bold] //bolded text
[i]Some Text[/i] //italicized text
[italics]Some Text[/italics] //italicized text
[u]Some Text[/u] //underlined text
[underline]Some Text[/underline] //underlined text
[s]some Text[/s] //strikedthrough text
[strike]Some Text[/strike] //strikethrough text
[code]{Some Code}[/code] //code text
[code]
Multiline code
Works a *little* bit differently
[i]And this should NEVER change[/i]
[/code]
[font=Verdana]Some Text[/font] //text in verdana font
[font size=20pt]Some Text[/font] //20pt size text
[font color=red]Some Text[/font] //red text. Must be a CSS color.
[font color=#000000]Some Text[/font] //text with the color of #000000. The # is necessary
[font variant=upper]Some Text[/font] //uppercased text
[font variant=lower]Some Text[/font] //lowercase text
[font variant=smallcaps]Some Text[/font] //smallcaps text
[size=20pt]Some Text[/size] //20pt size text
[color=red]Some Text[/color] //red text
[color=#000000]Some Text[/color] //text with the color of #000000. The # is necessary
[smallcaps]Some Text[/smallcaps] //smallcaps text
[url]https://darkstorm.techs[/url] //linked text
[url=https://darkstorm.tech]My website[/url] //linked text
[url title="My Website"]https://darkstorm.tech[/url] //linked text with title
[url tab]https://darkstorm.tech[/url] //link that opens into a new tab
[link]https://darkstorm.tech[/link] //linked text
[link=https://darkstorm.tech]My Website[/link] //linked text
[link title="Title"]https://darkstorm.tech[/link] //linked text with tooltip
[link tab]https://darkstorm.tech[/link] //link that opens into a new tab
[youtube]JsbdJFHRh6c[/youtube] //youtube video
[youtube height=200 width=500]JsbdJFHRh6c[/youtube] //youtube video with set size
[youtube=500x200]JsbdJFHRh6c[/youtube] //youtube video with set size
[youtube left]https://www.youtube.com/watch?v=JsbdJFHRh6c&t=6589s[/youtube] //youtube video floated left
[youtube right]https://www.youtube.com/live/JsbdJFHRh6c?si=K4_eqgHmXWIQTbTV[/youtube] //youtube video floated right
[img]test.png[/img] //an image
[img=500x200]test.png[/img] //an image with set size
[img height=200 width=500]test.png[/img] //an image with set size
[img left]test.png[/img] //an image floated left
[img right]test.png[/img] //an image floated right
[img alt="A simple test"]test.png[/img] //an image with alternate text
[img title="TEST TITLE"]test.png[/img] //an image with title
[image]test.png[/image] //same as [img] tag
[title]Some Text[/title] //Large text made for use as a title
[t1]Some Text[/t1] //Large text made for use as a title. Same as [title]
[t2]Some Text[/t2] //Slightly smaller text than [t1]. Meant for use as a title of some sort
[t3]Some Text[/t3] //Slightly smaller text than [t2]. Meant for use as a title of some sort
[t4]Some Text[/t4] //Slightly smaller text than [t3]. Meant for use as a title of some sort
[t5]Some Text[/t5] //Slightly smaller text than [t4]. Meant for use as a title of some sort
[t6]Some Text[/t6] //Slightly smaller text than [t5]. Meant for use as a title of some sort
[align=center]Some Text[/align] //Aligns the text. The text will be in a separate paragraph
[align center]Some Text[/align] //Equal sign is optional
[float=right]Floaty McFloat Face[/float] //Float the content content (for HTML, this will be a floated div)
[float right]Floaty McFloat Face[/float] //Equal sign is optional
[bullet]Bullet 1 * Bullet 2[/bullet] //bulleted list
[ul]
* Item 1
Item 2
[/ul] //an unordered (bulleted) list
[ol]
* Item 1
Item 2
[/ol] //an ordered (numbered) list
[bullet] * Item 1 * Item 2[/bullet] //same as
[number] * Item 1 * Item 2[/number] //same as
[ul]* Item 1 * Item 2[/ul] //an unordered (bulleted) list
[ol]* Item 1 * Item 2[/ol] //an ordered (numbered) list`
	bbTestResult = `<p><b>Some Text</b> //bolded text</p><p><b>Some Text</b> //bolded text</p><p><i>Some Text</i> //italicized text</p><p><i>Some Text</i> //italicized text</p><p>Some Text //underlined text</p><p>Some Text //underlined text</p><p><s>some Text</s> //strikedthrough text</p><p><s>Some Text</s> //strikethrough text</p><p><code>{Some Code}</code> //code text</p><p><pre><code>
Multiline code
Works a *little* bit differently
[i]And this should NEVER change[/i]
</code></pre></p><p><span style='font-family:Verdana;'>Some Text</span> //text in verdana font</p><p><span style='font-size:20pt;'>Some Text</span> //20pt size text</p><p><span style='color:red;'>Some Text</span> //red text. Must be a CSS color.</p><p><span style='color:#000000;'>Some Text</span> //text with the color of #000000. The # is necessary</p><p><span style='text-transform:uppercase;'>Some Text</span> //uppercased text</p><p><span style='text-transform:lowercase;'>Some Text</span> //lowercase text</p><p><span style='font-variant:small-caps;'>Some Text</span> //smallcaps text</p><p><span style='font-size:20pt'>Some Text</span> //20pt size text</p><p><span style='color:red'>Some Text</span> //red text</p><p><span style='color:#000000'>Some Text</span> //text with the color of #000000. The # is necessary</p><p><span style='font-variant:small-caps'>Some Text</span> //smallcaps text</p><p><a href='https://darkstorm.techs'>https://darkstorm.techs</a> //linked text</p><p><a href='https://darkstorm.tech'>My website</a> //linked text</p><p><a href='https://darkstorm.tech'title="My Website">https://darkstorm.tech</a> //linked text with title</p><p><a href='https://darkstorm.tech'target='_blank'>https://darkstorm.tech</a> //link that opens into a new tab (target=_blank)</p><p><a href='https://darkstorm.tech'>https://darkstorm.tech</a> //linked text</p><p><a href='https://darkstorm.tech'>My Website</a> //linked text</p><p><a href='https://darkstorm.tech'title="Title">https://darkstorm.tech</a> //linked text with tooltip</p><p><a href='https://darkstorm.tech'target='_blank'>https://darkstorm.tech</a> //link that opens into a new tab (target=_blank)</p><p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' allowfullscreen></iframe> //youtube video</p><p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='width:500px;height:200px;' allowfullscreen></iframe> //youtube video with set size</p><p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='width:500px;height:200px;' allowfullscreen></iframe> //youtube video with set size</p><p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='float:left;' allowfullscreen></iframe> //youtube video floated left</p><p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='float:right;' allowfullscreen></iframe> //youtube video floated right</p><p><img src='test.png'/> //an image</p><p><img src='test.png' style='width:500px;height:200px;'/> //an image with set size</p><p><img src='test.png' style='width:500px;height:200px;'/> //an image with set size</p><p><img src='test.png' style='float:left;'/> //an image floated left</p><p><img src='test.png' style='float:right;'/> //an image floated right</p><p><img src='test.png' alt="A simple test"/> //an image with alternate text</p><p><img src='test.png' title="TEST TITLE"/> //an image with title</p><p><img src='test.png'/> //same as [img] tag</p><p><h1>Some Text</h1> //Large text made for use as a title</p><p><h1>Some Text</h1> //Large text made for use as a title. Same as [title]</p><p><h2>Some Text</h2> //Slightly smaller text than [t1]. Meant for use as a title of some sort</p><p><h3>Some Text</h3> //Slightly smaller text than [t2]. Meant for use as a title of some sort</p><p><h4>Some Text</h4> //Slightly smaller text than [t3]. Meant for use as a title of some sort</p><p><h5>Some Text</h5> //Slightly smaller text than [t4]. Meant for use as a title of some sort</p><p><h6>Some Text</h6> //Slightly smaller text than [t5]. Meant for use as a title of some sort</p><p><div style='text-align:center;'>Some Text</div> //Aligns the text. The text will be in a separate paragraph</p><p><div style='text-align:center;'>Some Text</div> //Equal sign is optional</p><p><div style='float:right;'>Floaty McFloat Face</div> //Float the content content (for HTML, this will be a floated div)</p><p><div style='float:right;'>Floaty McFloat Face</div> //Equal sign is optional</p><p><ol><li>Bullet 1</li><li>Bullet 2</li></ol> //bulleted list</p><p><ul><li>Item 1</li><li>Item 2</li></ul> //an unordered (bulleted) list</p><p><ol><li>Item 1</li><li>Item 2</li></ol> //an ordered (numbered) list</p><p><ol><li>Item 1</li><li>Item 2</li></ol> //same as</p><p><ol><li>Item 1</li><li>Item 2</li></ol> //same as</p><p><ul><li>Item 1</li><li>Item 2</li></ul> //an unordered (bulleted) list</p><p><ol><li>Item 1</li><li>Item 2</li></ol> //an ordered (numbered) list</p>`
	mdTestString = ``
	mdTestResult = ``
)

func TestBBCode(t *testing.T) {
	conv := NewComboConverter()
	converted := conv.BBHTMLConvert(bbTestString)
	if converted != bbTestResult {
		t.Log("BB Conversion is:")
		t.Log(converted)
		t.Log("But should be:")
		t.Log(bbTestResult)
		t.Fatal("BB Conversion failed")
	}
	converted = conv.MarkdownHTMLConvert(mdTestString)
	if converted != mdTestResult {
		t.Log("BB Conversion is:")
		t.Log(converted)
		t.Log("But should be:")
		t.Log(bbTestResult)
		t.Fatal("BB Conversion failed")
	}
}
