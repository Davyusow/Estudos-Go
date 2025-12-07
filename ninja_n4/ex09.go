package main

import (
	"fmt"
)

func main() {
    pessoas := map[string][]string{
        "da_Silva_Beltrano" : {"Velejar","Pescar"},
        "Silveira_Fulano" : {"Ler","Estudar"},
        "Lima_Ciclano" : {"Correr","Malhar"},
    }

    for i,v := range pessoas {
        fmt.Println(i,v)
    }

    delete(pessoas,"Lima_Ciclano")
    println()

    for i,v := range pessoas {
        fmt.Println(i,v)
    }

}
