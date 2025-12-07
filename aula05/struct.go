package main

import "fmt"

type cliente struct {
    nome string
    sobrenome string
    fumante bool
}

func main() {
    cliente1 := cliente{
        nome: "Fulano",
        sobrenome: "Silva",
        fumante: true,
    }

    cliente2 := cliente{"Xavier", "X-men", false,}

    fmt.Println(cliente1)
    fmt.Println(cliente2)
}
