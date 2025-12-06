package main

import "fmt"

const (
    nome = 0
    sobrenome = 1
    hobby = 2
)

func main() {
    dados := [][]string {
        []string{
          "Fulano", //seria melhor botar nome, sobrenome e hobby tudo aqui pra cada 1
          "Beltrano",
          "Ciclano",
        },
        []string{
            "Da silva",
            "Silveira",
            "Da silva",
        },
        []string{
            "Pescar",
            "Dirigir",
            "Desenhar",
        },
    }

    for i := range(3) {
        fmt.Println(dados[nome][i],dados[sobrenome][i],"gosta de",dados[hobby][i])
    }
}
