package main

import (
	"flag"
	"fmt"

	"bbConvert"
)

func main() {
	wrap := flag.Bool("p", false, "Wrap each argument in <p> tags, if not present then will combine the arguments to form one large string before processing")
	classes := flag.String("class", "", "A list of classes for the external paragraph tags of each argument")
	flag.Parse()
	if len(flag.Args()) >= 1 {
		if *classes != "" {
			bbConvert.AddClass(*classes)
		}
		bbConvert.SetWrap(*wrap)
		var input string
		for _, v := range flag.Args() {
			input += v + "\n"
		}
		outFin := bbConvert.BBtohtml(input)
		fmt.Println(outFin)
	}
}
