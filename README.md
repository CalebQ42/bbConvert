#bbConvert
A package to convert bbcode to HTML.

#BBCode support
bbConvert has support for the following bb tags:

    [b]Some Text[/b] //bolded text
    [bold]Some Text[/bold] //bolded text
    [i]Some Text[/i] //italicized text
    [italics]Some Text[/italics] //italicized text
    [u]Some Text[/u] //underlined text
    [underline]Some Text[/underline] //underlined text
    [s]some Text[/s] //strikedthrough text
    [strike]Some Text[/strike] //strikethrough text
    [font=Verdana]Some Text[/font] //text in verdana font
    [font size=20pt]Some Text[/font] //20pt size text
    [font color=red]Some Text[/font] //red text
    [font color=#000000]Some Text[/font] //text with the color of #000000. The # is unnecessary
    [font variant=upper]Some Text[/font] //uppercased text
    [font variant=lower]Some Text[/font] //lowercase text
    [font variant=smallcaps]Some Text[/font] //smallcaps text
    [size=20pt]Some Text[/size] //20pt size text
    [color=red]Some Text[/color] //red text
    [color=#000000]Some Text[/color] //text with the color of #000000. The # is unnecessary
    [smallcaps]Some Text[/smallcaps] //smallcaps text
    [url]Link address[/url] //linked text
    [url=address]Some Text[/url] //linked text
    [url title="Title"]Link address[/url] //linked text with title
    [link]Link address[/link] //linked text
    [link=address]Some Text[/link] //linked text
    [link title="Title"]Link address[/link] //linked text with title
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
    [title]Some Text[/title] //Large text made for use as a title
    [t1]Some Text[/t1] //Large text made for use as a title. Same as [title]
    [t2]Some Text[/t2] //Slightly smaller text than [t1]. Meant for use as a title of some sort
    [t3]Some Text[/t3] //Slightly smaller text than [t2]. Meant for use as a title of some sort
    [t4]Some Text[/t4] //Slightly smaller text than [t3]. Meant for use as a title of some sort
    [t5]Some Text[/t5] //Slightly smaller text than [t4]. Meant for use as a title of some sort
    [t6]Some Text[/t6] //Slightly smaller text than [t5]. Meant for use as a title of some sort
    [align=center]Some Text[/align] //Aligns the insides (encapsulates the insides in a div)
    [ul]
    * Item 1
    * Item 2
    [/ul] //an unordered (bulleted) list
    [ol]
    * Item 1
    * Item 2
    [/ol] //an ordered (numbered) list
    [ul]* Item 1 * Item 2[/ul] //an unordered (bulleted) list
    [ol]* Item 1 * Item 2[/ol] //an ordered (numbered) list

If both img=/youtube= and height=/width= are present, img=/youtube= takes precedence.

If left unspecified then an img is set to width=20% and float=left (size is overridden when either height or width is set)

If left unspecified then youtube sets height=315 height=560

Both lists (ul, ol) trims spaces from the beginning and end of their items so spaces around * aren't necessary. If there is text before the first * it will be interpreted as it's own bullet/number

For the titles ([t1] - [t6]) then if you have a number less than 1 (such as zero) after the t then it will be the same as [t1] and if the number is greater than 6 it will be treated as [t6]

Tag and parameters aren't case sensitive unless they need to be (such as title and alt)

# Todo (Probably in order)

    [font align=center]Center, right, left, and justified support[/font]

# Known Issues
If "title=" is in the alt= then it could cause problems since it will try to parse the title from inside the alt
