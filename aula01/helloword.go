package main

import "fmt"

func main() {
    fmt.Println("Top",100,"Coisas Mais insanas para se fazer as 20:00","Horas")
    num_bytes, erros :=  fmt.Println("Essa frase tem 28 caractéres") // adiciona os 2 retornos da função nas variáveis
    // := cria a variável e define seu valor
    fmt.Println(num_bytes, erros)

    num_bytes, _ = fmt.Println("Eu só quero o tamanho!") //quando não se quer um dos retornos, colocar '_' para receber irá descartar ele imediatamente
    // = altera o valor da variável
    fmt.Println(num_bytes)

}
