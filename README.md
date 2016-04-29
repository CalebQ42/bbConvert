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
    [url=http://apage.com title="A Title"]link[/url]

If both img=/youtube= and height=/width= are present, img=/youtube= takes precedence.

If left unspecified then an img is set to width=20% and float=left (size is overridden when either height or width is set)

If left unspecified then youtube sets height=315 height=560

Both lists (ul, ol) trims spaces from the beginning and end of their items so spaces around * aren't necessary. If there is text before the first * it will be interpreted as it's own bullet/number

For the titles ([t1] - [t6]) then if you have a number less than 1 (such as zero) after the t then it will be the same as [t1] and if the number is greater than 6 it will be treated as [t6]

Tag and parameters aren't case sensitive (though parameter values are case senstive)

# Todo (Probably in order)

    [font size=20px]text[/font]
    [font color=red]red text[/font]
    [font color=#000000]the '#' before is necessary[/font]
    [font font-family=veranda]Veranda'd text[/font]

Make the # in front of a hex color code optional

Get rid of debugging Println's (These will come and go, but I'll probably forget to remove them often)

# Ideas
Make youtube video iframes and img's have a specific class so they can easily be formatted from css

# Example
Look at Test.go for the recommended way to implement this

# Known Issues
If title= is in the alt= then it could mess it will try to parse the title from inside the alt

if width=, height=, right, or left are present in alt AND title AND are placed in between the alt and title parameters then it will not be detected (I have an idea of how to fix it, but because it's so specific and unlikely it probably won't be fixed for a while)
