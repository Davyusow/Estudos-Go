package main

import "fmt"

type pessoa struct {
    nome string
    idade int
}

type profissional struct {
    conta pessoa
    titulo string
    salario float64
}

func main(){

    pessoa1 := pessoa{"Alfredo",30}

    pessoa2 := profissional{
        conta: pessoa{
            nome: "Beltrano",
            idade: 42,
        },
        titulo: "Empresário",
        salario: 2000.00,
    }

    fmt.Println(pessoa1)
    fmt.Println(pessoa2)
}
