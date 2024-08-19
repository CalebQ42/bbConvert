# bbConvert [![GoDoc](https://godoc.org/github.com/CalebQ42/bbConvert?status.svg)](https://godoc.org/github.com/CalebQ42/bbConvert) [![Coverage Status](https://coveralls.io/repos/github/CalebQ42/bbConvert/badge.svg?branch=master)](https://coveralls.io/github/CalebQ42/bbConvert?branch=master) [![Build Status](https://travis-ci.org/CalebQ42/bbConvert.svg?branch=master)](https://travis-ci.org/CalebQ42/bbConvert) [![Go Report Card](https://goreportcard.com/badge/github.com/CalebQ42/bbConvert)](https://goreportcard.com/report/github.com/CalebQ42/bbConvert)

bbConvert is an easy way to process BBCode and Markdown and convert it to HTML. BBCode additionally has support for custom processing to allow for custom tags or even conversion to something other then HTML

## Support

If you have any questions, feel free to open an issue and ask it. I don't have a life; I'll probably answer quickly.

## BBCode

```BBCode
[b]Some Text[/b] //bolded text
[bold]Some Text[/bold] //bolded text
[i]Some Text[/i] //italicized text
[italics]Some Text[/italics] //italicized text
[u]Some Text[/u] //underlined text
[underline]Some Text[/underline] //underlined text
[s]some Text[/s] //strikedthrough text
[strike]Some Text[/strike] //strikethrough text
[code]{Some Code}[/code] //code text
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
[url]Link address[/url] //linked text
[url=address]Some Text[/url] //linked text
[url title="Title"]Link address[/url] //linked text with title
[url tab]Link address[/url] //link that opens into a new tab (target=_blank)
[link]Link address[/link] //linked text
[link=address]Some Text[/link] //linked text
[link title="Title"]Link address[/link] //linked text with tooltip
[link tab]Link address[/link] //link that opens into a new tab (target=_blank)
[youtube]Youtube URL or video ID[/youtube] //youtube video
[youtube height=200 width=500]Youtube URL or video ID[/youtube] //youtube video with set size
[youtube=500x200]Youtube URL or video ID[/youtube] //youtube video with set size
[youtube left]Youtube URL or video ID[/youtube] //youtube video floated left
[youtube right]Youtube URL or video ID[/youtube] //youtube video floated right
[img]Image URL[/img] //an image
[img=500x200]Image URL[/img] //an image with set size
[img height=200 width=500]Image URL[/img] //an image with set size
[img left]Image URL[/img] //an image floated left
[img right]Image URL[/img] //an image floated right
[img alt="Alternate text"]Image URL[/img] //an image with alternate text
[img title="Title"]Image URL[/img] //an image with title
[image]Image URL[/image] //same as [img] tag
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
[bullet] * Item 1 * Item 2[/bullet] //same as [ul]
[number] * Item 1 * Item 2[/number] //same as [ol]
[ul]* Item 1 * Item 2[/ul] //an unordered (bulleted) list
[ol]* Item 1 * Item 2[/ol] //an ordered (numbered) list
```
