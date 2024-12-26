package bbConvert

import (
	"fmt"
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
	bbTestResult = `<p><b>Some Text</b> //bolded text</p>
<p><b>Some Text</b> //bolded text</p>
<p><i>Some Text</i> //italicized text</p>
<p><i>Some Text</i> //italicized text</p>
<p>Some Text //underlined text</p>
<p>Some Text //underlined text</p>
<p><s>some Text</s> //strikedthrough text</p>
<p><s>Some Text</s> //strikethrough text</p>
<p><code>{Some Code}</code> //code text</p>
<p><pre><code>
Multiline code
Works a *little* bit differently
[i]And this should NEVER change[/i]
</code></pre></p>
<p><span style='font-family:Verdana;'>Some Text</span> //text in verdana font</p>
<p><span style='font-size:20pt;'>Some Text</span> //20pt size text</p>
<p><span style='color:red;'>Some Text</span> //red text. Must be a CSS color.</p>
<p><span style='color:#000000;'>Some Text</span> //text with the color of #000000. The # is necessary</p>
<p><span style='text-transform:uppercase;'>Some Text</span> //uppercased text</p>
<p><span style='text-transform:lowercase;'>Some Text</span> //lowercase text</p>
<p><span style='font-variant:small-caps;'>Some Text</span> //smallcaps text</p>
<p><span style='font-size:20pt'>Some Text</span> //20pt size text</p>
<p><span style='color:red'>Some Text</span> //red text</p>
<p><span style='color:#000000'>Some Text</span> //text with the color of #000000. The # is necessary</p>
<p><span style='font-variant:small-caps'>Some Text</span> //smallcaps text</p>
<p><a href='https://darkstorm.techs'>https://darkstorm.techs</a> //linked text</p>
<p><a href='https://darkstorm.tech'>My website</a> //linked text</p>
<p><a href='https://darkstorm.tech'title="My Website">https://darkstorm.tech</a> //linked text with title</p>
<p><a href='https://darkstorm.tech'target='_blank'>https://darkstorm.tech</a> //link that opens into a new tab</p>
<p><a href='https://darkstorm.tech'>https://darkstorm.tech</a> //linked text</p>
<p><a href='https://darkstorm.tech'>My Website</a> //linked text</p>
<p><a href='https://darkstorm.tech'title="Title">https://darkstorm.tech</a> //linked text with tooltip</p>
<p><a href='https://darkstorm.tech'target='_blank'>https://darkstorm.tech</a> //link that opens into a new tab</p>
<p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' allowfullscreen></iframe> //youtube video</p>
<p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='width:500px;height:200px;' allowfullscreen></iframe> //youtube video with set size</p>
<p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='width:500px;height:200px;' allowfullscreen></iframe> //youtube video with set size</p>
<p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='float:left;' allowfullscreen></iframe> //youtube video floated left</p>
<p><iframe src='https://youtube.com/embed/JsbdJFHRh6c' style='float:right;' allowfullscreen></iframe> //youtube video floated right</p>
<p><img src='test.png'/> //an image</p>
<p><img src='test.png' style='width:500px;height:200px;'/> //an image with set size</p>
<p><img src='test.png' style='width:500px;height:200px;'/> //an image with set size</p>
<p><img src='test.png' style='float:left;'/> //an image floated left</p>
<p><img src='test.png' style='float:right;'/> //an image floated right</p>
<p><img src='test.png' alt="A simple test"/> //an image with alternate text</p>
<p><img src='test.png' title="TEST TITLE"/> //an image with title</p>
<p><img src='test.png'/> //same as [img] tag</p>
<p><h1>Some Text</h1> //Large text made for use as a title</p>
<p><h1>Some Text</h1> //Large text made for use as a title. Same as [title]</p>
<p><h2>Some Text</h2> //Slightly smaller text than [t1]. Meant for use as a title of some sort</p>
<p><h3>Some Text</h3> //Slightly smaller text than [t2]. Meant for use as a title of some sort</p>
<p><h4>Some Text</h4> //Slightly smaller text than [t3]. Meant for use as a title of some sort</p>
<p><h5>Some Text</h5> //Slightly smaller text than [t4]. Meant for use as a title of some sort</p>
<p><h6>Some Text</h6> //Slightly smaller text than [t5]. Meant for use as a title of some sort</p>
<p><div style='text-align:center;'>Some Text</div> //Aligns the text. The text will be in a separate paragraph</p>
<p><div style='text-align:center;'>Some Text</div> //Equal sign is optional</p>
<p><div style='float:right;'>Floaty McFloat Face</div> //Float the content content (for HTML, this will be a floated div)</p>
<p><div style='float:right;'>Floaty McFloat Face</div> //Equal sign is optional</p>
<p><ol><li>Bullet 1</li><li>Bullet 2</li></ol> //bulleted list</p>
<p><ul><li>Item 1</li><li>Item 2</li></ul> //an unordered (bulleted) list</p>
<p><ol><li>Item 1</li><li>Item 2</li></ol> //an ordered (numbered) list</p>
<p><ol><li>Item 1</li><li>Item 2</li></ol> //same as</p>
<p><ol><li>Item 1</li><li>Item 2</li></ol> //same as</p>
<p><ul><li>Item 1</li><li>Item 2</li></ul> //an unordered (bulleted) list</p>
<p><ol><li>Item 1</li><li>Item 2</li></ol> //an ordered (numbered) list</p>`
	//TODO
	mdTestString = "```\nThis is some code that\n*should not*\nGet ***converted***\n```\n\nCode also comes in an `inline variation`\n" + `
# Markdown test

## Bullet test

* This is a test of the bullet points
* And if it can handle *formatting **within** the bullet*
  * And of course multiple _levels_
	* of bullets
    1) ~~Can it handle mixed? I don't think so, not yet~~ DONE

### Numbered list test

1) Of course we can't forget __numbered lists__
2) Where we can use ) or
3. dots. And of course
1000) ***We don't actually care what number it is***
  1) And should have multi-level support
  2) *just like bullets*

#### And don't forget block quotes

> This is a quote with multiple lines
>
> And junk
>> And a nested quote

##### Link test

Let's not forget about [links](https://darkstorm.tech) and, of course, images:

![test image](test.png)
`
	mdTestResult = `<p><pre><code>
This is some code that
*should not*
Get ***converted***
</code></pre></p>
<p>Code also comes in an <code>inline variation</code></p>
<p><h1>Markdown test</h1></p>
<p><h2>Bullet test</h2></p>
<p><ul><li>This is a test of the bullet points</li><li>And if it can handle <i>formatting <b>within</b> the bullet</i></li><ul><li>And of course multiple <i>levels</i></li><li>of bullets</li><ol><li><s>Can it handle mixed? I don't think so, not yet</s> DONE</li></ol></ul></ul>
<h3>Numbered list test</h3></p>
<p><ol><li>Of course we can't forget <b>numbered lists</b></li><li>Where we can use ) or</li><li>dots. And of course</li><li><b><i>We don't actually care what number it is</i></b></li><ol><li>And should have multi-level support</li><li><i>just like bullets</i></li></ol></ol>
<h4>And don't forget block quotes</h4></p>
<p><blockquote><p>This is a quote with multiple lines</p><p>And junk</p><blockquote><p>And a nested quote</p></blockquote></p></blockquote>
<h5>Link test</h5></p>
<p>Let's not forget about <a href='https://darkstorm.tech'>links</a> and, of course, images:</p>
<p><img src='test.png' alt='test image'>
</p>`
	testy = `I've recently gotten into the bad habit of looking at software dev twitter (no I'm never calling it X) and have been constantly annoyed at the amount of artificial benchmarks people share. The latest one to draw my ire (and spawn this post) is a _bad_ "benchmark" that's basically just 1 BILLION iterations of a for loop.

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">More languages, more insights!<br><br>A few interesting takeaways:<br><br>* Java and Kotlin are quick! Possible explanation: Google is heavily invested in performance here.<br>* Js is really fast as far as interpreted / jit languages go.<br>* Python is quite slow without things like PyPy. <a href="https://t.co/GIshus2UXO">pic.twitter.com/GIshus2UXO</a></p>— Ben Dicken (@BenjDicken) <a href="https://twitter.com/BenjDicken/status/1861072804239847914?ref_src=twsrc%5Etfw">November 25, 2024</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

Now I cannot talk to most of the languages shown, but I have significant experience in Go and have spent a not insignificant time optimizing Go code (in particular my squashfs library). The second I opened up the code for this "benchmark" I knew that whoever had written this code has never tried to write optimized Go code. First let's start with the results without any changes. For simplicity I'll only show the results of C and Go.

` + "```" + `
C = 1.29s
C = 1.29s
C = 1.29s

Go = 1.51s
Go = 1.51s
Go = 1.51s
` + "```" + `

This is fairly expected, as it's what's in line with the post and what is logical, Go's structure is fairly low level and similar to C, but it is garbage compiled meaning it _will_ be slower in real world applications. Now let's look at the results of my optimized code:

` + "```" + `
C = 1.29s
C = 1.29s
C = 1.29s

Go = 1.29s
Go = 1.30s
Go = 1.29s
` + "```" + `

Suddenly, C's lead is gone! _What black magic is this???_. Well, if you actually look at the original code and you know Go, you'll probably notice it immediately: the "benchmark" is using ` + "`" + `int` + "`" + `. That's right, my optimizations boiled down to making all ` + "`" + `int` + "`" + ` instances ` + "`" + `int32` + "`" + `s. I'm honestly a bit surprised it basically ties C, but I suspect that, _since this isn't a real world benchmark_, the garbage collector never actually has to do anything, meaning Go's primary disadvantage is non-existent.

## My gaps in knowledge

Let me be clear, I am no expert, I do not actually know _why_ ` + "`" + `int32` + "`" + ` is faster then ` + "`" + `int` + "`" + `, I just know it is (I have theories, but that's all they are). Though I know many of the other languages, I haven't ever done any research on how to optimize them. It's possible all the other languages are perfectly optimized, but the fact such a simple optimization was overlooked invalidates the entire test in my mind.

## The Point

Let me be clear, benchmarks are important and useful, but the most useful benchmarks I've seen are between code of the same language as it removes a lot of the compiler magic and skill issues. Funnily enough, my benchmark between the code using ` + "`" + `int` + "`" + ` and ` + "`" + `int32` + "`" + ` _is_ a useful benchmark. The problem arises when you try to benchmark between fundamentally different languages (or even frameworks), but do not give them _all_ the same amount of time and attention. As an example, if I were to write the C code for this test we'd probably see Go with a lead, not because Go is faster, but because I know how to write optimized Go.

The real world is messy, and between DB calls, API requests, and IO, the actual performance gains/failure of any particular language becomes a lot more complex and their performance will largely depend on your needs. The vast majority of the time spending time optimizing code would be far better then re-writing in a different language. The only time I'd actually recommend switching languages is when you've already optimized and are still running into performance constraints __or__ if you want to learn. Let me be clear: [Ben Dicken](https://x.com/BenjDicken) _is_ a better engineer then me, but that doesn't mean he can't make mistakes.`
)

func TestBBCode(t *testing.T) {
	conv := NewComboConverter()
	converted := conv.BBHTMLConvert(bbTestString)
	if converted != bbTestResult {
		fmt.Print("BB Conversion should be:\n\n")
		fmt.Println(bbTestResult)
		fmt.Print("\nBut is:\n\n")
		fmt.Println(converted)
		t.Fatal("BB Conversion failed")
	}
	converted = conv.MarkdownHTMLConvert(mdTestString)
	if converted != mdTestResult {
		fmt.Print("Markdown Conversion should be:\n\n")
		fmt.Println(mdTestResult)
		fmt.Print("\nBut is:\n\n")
		fmt.Println(converted)
		t.Fatal("Markdown Conversion failed")
	}
}

func TestTesty(t *testing.T) {
	conv := NewComboConverter()
	converted := conv.MarkdownHTMLConvert(testy)
	t.Fatal(converted)
}
