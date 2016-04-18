# BBConverter
Converter from BBCode to HTML
# Usage
It takes in a single argument that is the string that has BB in it and will print (not println) it out to the standard console. A simple way to implement it in PHP is:
'''php
$in = escapeshellargs($in);
echo shell_exec("./BBConverter ".$in);
'''
It works and currently has support for:
[b]
[i]
[img]http://apage.com/image.png[/img]
[url]http://apage.com[/url]
[url=http://apage.com]link[/here]
