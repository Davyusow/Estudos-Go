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

    // fmt.Println(pessoas)

    sabores := map[string][]string{
        pessoas[0].sobrenome : pessoas[0].sabores,
        pessoas[1].sobrenome : pessoas[1].sabores,
    }

    for i := range pessoas {
        fmt.Println(pessoas[i].nome,pessoas[i].sobrenome)
        for j := range sabores {
            fmt.Println(sabores[j]) //sei lá vei to com sono algo assim
        }
    }

    fmt.Println(sabores["Silveira"])
}
