package bbConvert

import (
	"fmt"
	"testing"
)

func TestStuff(t *testing.T) {
	str := `[b]Hello[/b] [i]TADA[/i] [font=Verdana]VERDANA[/font] [t1]hello[/t1] [title]hello again[/title] [t6]yo[/t6] [t7]nothing to see here[/t7] [font=\"Noto Sans\" color=red variant=\"this is a test\"]Noto AND red[/font]`
	conv := NewBBConverter()
	fmt.Println(conv.HTMLConvert(str))
	t.Fatal("hi")
}
