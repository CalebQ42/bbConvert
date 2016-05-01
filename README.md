# BBConverter
Converter from BBCode to HTML
# Usage
A Simple function to convert BBCode to HTML

It currently has support for:

    [b]bold[/b]
    [i]italic[/i]
    [u]underline[/u]
    [s]strikethrough[/s]
    [color=red]The color must be html compatable[/color]
    [color=#000000]The '#' is optional[/color]
    [font=verdana]Verdana'd text[/font]
    [font size=20pt]text[/font]
    [font color=red]red text[/font]
    [font color=#000000]the '#' is optional[/font]
    [font variant=smallcaps]font in small caps[/font]
    [font variant=upper]FONT IN UPPER[/font]
    [font variant=lower]font in lower[/font]
    [size=12]12pt size text[/size]
    [smallcaps]This is in small capps :)[/smallcaps]
    [img]image URL[/img]
    [url]http://apage.com[/url]
    [url=http://apage.com]link[/here]
    [url=http://apage.com title="A Title"]link[/url]
    [url title="A Title"]http://apage.com[/url]
    [img=500x200]http://apage.com/image.png[/img]
    [img height=200 width=500]http://apage.com/image.png[/img]
    [img left]This image is floated left[/img]
    [img right]This image is floated right[/img]
    [img alt="Alternate for if picture doesn't show up"]image URL[/img]
    [img title="This shows up when hovering over the picture"]image URL[/img]
    [youtube]https://www.youtube.com/watch?v=U-G4TZzVeZ0[/youtube]
    [youtube]https://youtu.be/U-G4TZzVeZ0[/youtube]
    [youtube]U-G4TZzVeZ0[/youtube]
    [youtube height=200 width=500]Youtube URL or ID[/youtube]
    [youtube=500x200]Youtube URL[/youtube]
    [youtube left]U-G4TZzVeZ0[/youtube]
    [youtube right]U-G4TZzVeZ0[/youtube]
    [ul]
    * Item 1
    * Item 2
    [/ul]
    [ol]
    * Item 1
    * Item 2
    [/ol]
    [ul]* Item 1 * Item 2[/ul]
    [ol]* Item 1 * Item 2[/ol]
    [title]This will be big[/title]
    [t1]Same as title[/t1]
    [t2]Smaller than t1[/t2]
    [t3]Smaller than t2[/t3]
    [t4]Smaller than t3[/t4]
    [t5]Smaller than t4[/t5]
    [t6]Smaller than t5[/t6]

If both img=/youtube= and height=/width= are present, img=/youtube= takes precedence.

If left unspecified then an img is set to width=20% and float=left (size is overridden when either height or width is set)

If left unspecified then youtube sets height=315 height=560

Both lists (ul, ol) trims spaces from the beginning and end of their items so spaces around * aren't necessary. If there is text before the first * it will be interpreted as it's own bullet/number

For the titles ([t1] - [t6]) then if you have a number less than 1 (such as zero) after the t then it will be the same as [t1] and if the number is greater than 6 it will be treated as [t6]

Tag and parameters aren't case sensitive (though parameter values are case senstive)

# Todo (Probably in order)

    [font align=center]Center, right, left, and justified support[/font]
    [align=center]Center, right, left, and justified support[/align]
    [ul]
    * bullet
        * Space four times for a sub bullet [/ul]

Organize supported bbcode list and create github wiki (maybe)

# Ideas
Make youtube video iframes and img's have a specific class so they can easily be formatted from css

If there's any other bb code you think should be added, PLEASE TELL ME

# Example
Look at Test.go for the recommended way to implement this

If you put in something like:

    [title][color=blue]This is an example[/color][/title]
    [t3]This was actually parsed by my test program[/t3]
    [t4]Some features[/t4]
    [ul] *[i]Works quickly and with relatively few resources[/i]
    * [b]Writen in Go! so it is cross platform[/b]
    * [smallcaps]Has support for wrapping the output in paragraph tags[/smallcaps]
    * You can write HTML right into the <b>BB</b>
    * [font color=red variant=upper size=20pt]Made to work with a multitude[/font] of BB tags[/color]
    * [B][ColOR=009900]The bb tags aren't even case senstive[/color][/b][/ul]
    [u]In general, this is made to be extremely flexible and can be easily used on a server.[/u]
    Unfortunately making newlines a different paragraph isn't supported in the program, but a simple php script that explodes and implodes does the trick.

### You get something like:
    <h1><span style='color:blue;'>This is an example</span></h1><h3>This was actually parsed by my test program</h3><h4>Some features</h4><p><ul><li><i>Works quickly and with relatively few resources</i></li><li><b>Writen in Go! so it is cross platform</b></li><li><span style='font-variant:small-caps;'>Has support for wrapping the output in paragraph tags</span></li><li>You can write HTML right into the <b>BB</b></li><li><span style="font-size:20pt;color:red;">MADE TO WORK WITH A MULTITUDE</span> of BB tags[/color]</li><li><B><span style='color:009900;'>The bb tags aren't even case senstive</span></B></li></ul></p><p><u>In general, this is made to be extremely flexible and can be easily used on a server.</u></p><p>Unfortunately making newlines a different paragraph isn't supported in the program, but a simple php script that explodes and implodes does the trick.</p>

<h1><span style='color:blue;'>This is an example</span></h1><h3>This was actually parsed by my test program</h3><h4>Some features</h4><p><ul><li><i>Works quickly and with relatively few resources</i></li><li><b>Writen in Go! so it is cross platform</b></li><li><span style='font-variant:small-caps;'>Has support for wrapping the output in paragraph tags</span></li><li>You can write HTML right into the <b>BB</b></li><li><span style="font-size:20pt;color:red;">MADE TO WORK WITH A MULTITUDE</span> of BB tags[/color]</li><li><B><span style='color:009900;'>The bb tags aren't even case senstive</span></B></li></ul></p><p><u>In general, this is made to be extremely flexible and can be easily used on a server.</u></p><p>Unfortunately making newlines a different paragraph isn't supported in the program, but a simple php script that explodes and implodes does the trick.</p>

# Known Issues
If "title=" is in the alt= then it could cause problems since it will try to parse the title from inside the alt
