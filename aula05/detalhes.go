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

    fmt.Println("Pessoa1, nome:",pessoa1.nome)
    fmt.Println("Pessoa1, idade:",pessoa1.idade)

    fmt.Println("Pessoa2, nome:",pessoa2.conta.nome)


}
