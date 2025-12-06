package main

import "fmt"

func main() {

    amigos := map[string]int{
        "fulano": 213,
        "beltrano": 182,
        "fulana": 251,
        "ciclana": 987,
    }
    fmt.Println(amigos)
    fmt.Println("Ciclana: ",amigos["ciclana"])

    sera, ok := amigos["fulano"] // o 'ok' serve para receber um booleano retornando true caso exista aquele valor mapeado
    fmt.Println(sera,ok) //comma ok

    //Exemplo que me veio utilizando map e ok
    // inclusive, map é uma estrutura de dados muito perfomática
    nome := "beltrano"
    if login, ok := amigos[nome]; ok {
        fmt.Printf("Bem vindo!, %s\ntoken: %v\n\n",nome, login);
    } else {
        fmt.Println("Recusado!")
    }

    for chave, valor := range amigos {
        fmt.Println(chave,valor) //imprime na ordem de declaração
    }

    delete(amigos, "fulano") // deleta pela chave
    println()

    for chave, valor := range amigos {
        fmt.Println(chave,valor) //imprime na ordem de declaração
    }

}
