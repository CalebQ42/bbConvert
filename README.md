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
    [img]http://apage.com/image.png[/img]
    [url]http://apage.com[/url]
    [url=http://apage.com]link[/here]
    [img=500x200]http://apage.com/image.png[/img]
    [img height=200 width=500]http://apage.com/image.png[/img]
    
#Todo

    [img left]This image is floated left[/img]
    [img right]This image is floated right[/img]
    [img alt="Alternate for if picture doesn't show up"]image URL[/img]
    [img title="This shows up when hovering over the picture"]image URL[/img]

If both img= and height=/width= are present, height=/width= takes precedence
More is to come
