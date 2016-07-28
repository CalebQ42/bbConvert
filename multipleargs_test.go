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
