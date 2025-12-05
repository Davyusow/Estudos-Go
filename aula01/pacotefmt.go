package main

import "fmt"

func main(){
    x := "Ola como vai!\nEsse texto\n\tÉ um exemplo de um String interpretada"
    y := `
    Olá! Este é um exemplo
    de uma String Crua!
    Usando \n Cráse`
    fmt.Println(x)
    fmt.Println(y) //cada caractere é um 'rune literal'

    texto := "Olá super texto"
    fmt.Print(texto) // imprime sem pular linha
    fmt.Println(texto) // imprime pulando linha
    fmt.Printf("\t\n%s\n",texto) //imprime com a possibilidade de usar formatações

    texto2 := fmt.Sprint(texto + texto) //o Sprint retorna uma string, o + concatena as Strings
    fmt.Println(texto2)
    texto2 = fmt.Sprintln(texto + texto)
    fmt.Println(texto2)
    texto2 = fmt.Sprintf(texto, 2, 3,"\n")
    fmt.Println(texto2)
}
