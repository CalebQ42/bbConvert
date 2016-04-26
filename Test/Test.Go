package main

import(
    "flag"
    "fmt"
    "github.com/CalebQ42/bbConvert"
)

func main(){
    wrap := flag.Bool("p",false,"Wrap each argument in <p> tags, if not present then will combine the arguments to form one large string before processing")
    classes := flag.String("class","","A list of classes for the external paragraph tags of each argument")
    style := flag.String("style","","The style for the external paragraph tags")
    flag.Parse()
    if len(flag.Args())>=1{
        if *classes != ""{
            bbConvert.AddParagraphClass(*classes)
        }
        if *style != ""{
            bbConvert.AddStyle(*style)
        }
        outFin := bbConvert.Convert(flag.Args(),*wrap)
        fmt.Println(outFin)
    }
}