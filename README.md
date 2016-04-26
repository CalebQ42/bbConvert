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
    [color=#000000]The '#' is necessary if doing it like this[/color]
    [img]image URL[/img]
    [url]http://apage.com[/url]
    [url=http://apage.com]link[/here]
    [img=500x200]http://apage.com/image.png[/img]
    [img height=200 width=500]http://apage.com/image.png[/img]
    [img left]This image is floated left[/img]
    [img right]This image is floated right[/img]
    [img alt="Alternate for if picture doesn't show up"]image URL[/img]
    [img title="This shows up when hovering over the picture"]image URL[/img]
    [youtube]https://www.youtube.com/watch?v=U-G4TZzVeZ0[/youtube]
    [youtube]https://youtu.be/U-G4TZzVeZ0[/youtube]
    [youtube]U-G4TZzVeZ0[/youtube]
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
    
If both img= and height=/width= are present, img= takes precedence. Support for more BB is coming (See below).

If left unspecified then an img is set to width=20% and float=left (size is overridden when either height or width is set)

If left unspecified then youtube sets height=315 height=560

Both lists (ul, ol) trims spaces from the beginning and end of their items so spaces around * aren't necessary. If there is text before the first * it will be interpreted as it's own bullet/number

Tag and parameters aren't case sensitive (though parameter values are case senstive)

#Todo (Probably in order)

    [title]This will be big[/title]
    [t1]Same as title[/t1]
    [t2]Smaller than t1[/t2]
    [youtube height=200 width=500]Youtube URL[/youtube]
    [youtube=500x200]Youtube URL[/youtube]
    [youtube left]U-G4TZzVeZ0[/youtube]
    [youtube right]U-G4TZzVeZ0[/youtube]
    [url=http://apage.com title="A Title"]link[/url]
    [font size=20px]text[/font]
    [font color=red]red text[/font]
    [font color=#000000]the '#' before is necessary[/font]
    [font font-family=veranda]Veranda'd text[/font]

Make the # in front of a hex color code optional

Get rid of debugging Println's (These will come and go, but I'll probably forget to remove them often)

Go fmt the code (I'm just being lazy right now)

#Ideas

Make it so that tags that are a t followed by a number automagically sets it's font size, allowing for rediculously small text (like [t500])

Make youtube video iframes and img's have a specific class so they can easily be formatted from css

#Example
Look at Test.go for the recommended way to implement this

#Known Issues
if width=, height=, right, or left are present in alt AND title AND are placed inbetween the alt and title parameters then it will not be detected (I have an idea of how to fix it, but because it's so specific and unlikely it probably won't be fixed for a while)
