package main

import (
	"fmt"
)

func main() {
    s := "\tHello, world!" // o String é interpretado
    texto := `s ada  \nsdasd` //string cru

    sb := []byte(s)

    fmt.Printf("%v\n%T\n", s, s)
    fmt.Printf("%v\n\n",texto)
    fmt.Printf("%v\n%T\n",sb,sb)

    for _, v := range sb{   // foreach
        fmt.Printf("%v- %T - %#U - %#x\n", v,v,v,v)
    }

    println()
    for i := 0; i < len(s); i++ {   // for padrão
        fmt.Printf("%v- %T - %#U - %#x\n", s[i],s[i],s[i],s[i])
    }

}
