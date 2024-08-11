package bbConvert

import (
	"fmt"
	"testing"
)

func TestStuff(t *testing.T) {
	str := "[b]Hello[/b] [i]TADA[/i] [font=Verdana]VERDANA[/font] [t1]hello[/t1] [title]hello again[/title] [t6]yo[/t6] [t7]nothing to see here[/t7]"
	conv := NewBBConverter()
	fmt.Println(conv.Convert(str))
	t.Fatal("hi")
}
