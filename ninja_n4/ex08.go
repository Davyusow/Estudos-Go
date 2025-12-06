package main

import (
	"fmt"
)

func main() {
    pessoas := map[string][]string{
        "Beltrano" : {"da Silva","Pescar"},
        "Fulano" : {"Silveira","Estudar"},
        "Ciclano" : {"Lima","Malhar"},
    }

    nome := "Beltrano"
    fmt.Println(nome,pessoas[nome])
}
