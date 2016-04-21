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
    
If both img= and height=/width= are present, height=/width= takes precedence. Support for more BB is coming (probably).

#Todo

    [youtube]Youtube URL[/youtube]
    [url=http://apage.com title="A Title"]link[/url]

Get rid of debugging Println's
Go fmt the code (I'm just being lazy right now)

#Example
Look at Test.go for an easy way to implement this (it  even uses concurrency :) ) with support for multiple arguments and the option to wrap the arguments in paragraph tags.

#Known Issues
if width=, height=, right, or left are present in alt AND title AND are placed inbetween the alt and title parameters then it will not be detected (I have an idea of how to fix it, but because it's so specific and unlikely it probably won't be fixed for a while)