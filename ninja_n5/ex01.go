package main

import "fmt"

type pessoa struct {
    nome string
    sobrenome string
    sabores []string
}

func main() {
    pessoas := make([]pessoa,2)

    pessoas[0] = pessoa{
        nome: "Fulano",
        sobrenome: "da Silva",
        sabores: []string{"chocolate", "morango"},
    }

    pessoas[1] = pessoa{
        nome: "Beltrano",
        sobrenome: "Silveira",
        sabores: []string{"baunilha", "menta"},
    }

    fmt.Println(pessoas)
}
