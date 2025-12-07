package main

import "fmt"

type pessoa struct {
    nome string
    idade int
}

func (p pessoa) atualizar(nome string, idade int) pessoa {
    p.nome = nome
    p.idade = idade
    return  p
}

func (p pessoa) imprimir() {
    fmt.Println("Nome: ",p.nome)
    fmt.Println("Idade: ",p.idade)
}

func main() {
    fulano := pessoa{
        "Fulano",
        12,
    }

    fulano.imprimir()

    fulano = fulano.atualizar("Fulano",13)

    fulano.imprimir()


}
