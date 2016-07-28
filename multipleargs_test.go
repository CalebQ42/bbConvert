package bbConvert

import "testing"

func TestMultArgsImg(t *testing.T) {
	var exarr []string
	in := "[img=500x200 title='Not an actual URL' alt='This will always display with this URL']http://aplace.com/image.png[/img]"
	expected := "<p><img style='float:none;width:500;height:200;' alt='This will always display with this URL' title='Not an actual URL' src='http://aplace.com/image.png'/></p>"
	exarr = append(exarr, expected)
	expected = "<p><img style='width:500;height:200;float:none;' alt='This will always display with this URL' title='Not an actual URL' src='http://aplace.com/image.png'/></p>"
	exarr = append(exarr, expected)
	expected = "<p><img style='height:200;float:none;width:500;' alt='This will always display with this URL' title='Not an actual URL' src='http://aplace.com/image.png'/></p>"
	exarr = append(exarr, expected)
	conv := CreateConverter(true, true)
	out := conv.Convert(in)
	for _, v := range exarr {
		if out == v {
			return
		}
	}
	t.Fail()
}

func TestLists(t *testing.T) {
	testStr := "[ul]\n* Level 1[ol]*Level2[ul]Level3[ul]\nLevel4[ol]Level 5[ol]* Level 5[/ol][/ol][/ul][/ul][/ol][/ul]"
	expected := "<ul><li>Level 1<ol><li>Level2<ul><li>Level3<ul><li>Level4<ol><li>Level 5<ol><li>Level 5</li></ol></li><li>/li></ol></li><li>/li></ul></li><li>/li></ul></li><li>/li></ol></li></ul>"
	conv := CreateConverter(true, true)
	out := conv.Convert(testStr)
	if out != expected {
		t.Fatal("Got: " + out + ", Expected: " + expected)
	}
}
