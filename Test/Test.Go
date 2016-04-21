package main

import(
    "flag"
    "fmt"
    "github.com/CalebQ42/bbConvert"
)

var wrap *bool

func main(){
    wrap = flag.Bool("p",false,"Wrap each argument in <p> tags, if not present then will combine the arguments to form one large string before processing")
    flag.Parse()
    if len(flag.Args())>=1{
        ch := make(chan string,len(flag.Args()))
        out := make([]string,len(flag.Args()))
        outFin := ""
        for _,v := range flag.Args(){
            go multiThreading(ch,v)
        }
        for range out{
            outFin += <-ch
        }
        fmt.Println(outFin)
    }
}

func multiThreading(ch chan string, in string){
    out := bbConvert.Convert(in)
    if *wrap{
        out = "<p>" + out + "</p>"
    }
    ch <- out
}