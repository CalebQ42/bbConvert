package main

import(
    "flag"
    "fmt"
    "github.com/CalebQ42/bbConvert"
)

func main(){
    flag.Parse()
    if len(flag.Args()==1){
        fmt.Println(bbConvert.Convert(flag.Args()[0]))
    }
}