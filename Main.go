package main

import(
    "flag"
    "fmt"
    "strings"
)

func main(){
    flag.Parse()
    in := flag.Args()[0]
    out := convert(in)
    fmt.Print(out)
}

func convert(str string) string{
    nd := -1
    for i:=0;i<len(str);i++{
        v := str[i]
        if v == '['{
            for j:= len(str)-1;j>-1;j--{
                if str[j]==']'{
                    nd = j
                    break
                }
            }
            if nd!=-1{
                tmp := toHTML(str[i:nd+1])
                if str[i:nd+1]!=tmp{
                    str = str[:i] + tmp + str[nd+1:]
                    i = 0
                    nd = -1
                }
            }
        }
    }
    return str
}

func toHTML(str string) string{
    var beg,end string
    for i,v := range str{
        if v ==']' || v==' ' || v=='='{
            beg = str[1:i]
            break
        }
    }
    var tmp string
    for i:=len(str)-3;i>0;i--{
        if str[i:i+2]=="[/"{
            tmp = str[i:]
            end = str[i+2:len(str)-1]
            break;
        }
    }
    if beg != end{
        return str
    }
    if beg =="url"||beg=="quote"||beg=="color"||beg=="style"{
        for i,v := range str{
            if v ==']'{
                beg = str[1:i]
                break
            }
        }
    }
    if strings.HasPrefix(tmp,"[/") && strings.HasSuffix(tmp,"]") && !isBBTag(tmp[2:len(tmp)-1]){
        return str
    }
    if len(str) - len(tmp) >1{
        str = "[" + convert(str[1:len(str)-len(tmp)]) + tmp
    }
    str = bbToTag(str,beg)
    return str
}

func isBBTag(str string) bool{
    str = strings.ToLower(str)
    tf := str=="b"||str=="i"||str=="u"||str=="s"||str=="url"||str=="img"||str=="quote"||str=="style"||str=="color"
    return tf
}

func bbToTag(str,bb string) string{
    if bb=="img"{
        str = "<img style='float:left;width:20%;' src='" + str[5:len(str)-6] + "'/>"
    }else if bb=="b" || bb=="i" || bb=="u" || bb=="s"{
        str = strings.Replace(str[:4],"[","<",1) + str[4:]
        str = strings.Replace(str[:4],"]",">",1) + str[4:]
        str = str[:len(str)-4] + strings.Replace(str[len(str)-4:],"[","<",1)
        str = str[:len(str)-4] + strings.Replace(str[len(str)-4:],"]",">",1)
    }else if bb=="url"{
        str = "<a href='" + str[5:len(str)-6] + "'>" + str[5:len(str)-6] + "</a>"
    }else if strings.HasPrefix(bb,"url="){
        str = "<a href='" + bb[5:]+"'>" + str[len(bb)+2:len(str)-6] + "</a>"
    }else if strings.HasPrefix(bb,"color="){
        str = "<span style='color:" + bb[7:] + ";'>" + str[len(bb)+2:len(str)-8] + "</span>"
    }else if strings.HasPrefix(bb,"quote=\"")|| strings.HasPrefix(bb,"quote='"){
        str = "<div class='quote'>"+bb[7:len(bb)-1]+"<blockquote>"+str[len(bb)+2:len(str)-8]+"</blockquote></div>"
    }else if strings.HasPrefix(bb,"quote="){
        str = "<div class='quote'>"+bb[6:]+"<blockquote>"+str[len(bb)+2:len(str)-8]+"</blockquote></div>"
    }
    return str
}