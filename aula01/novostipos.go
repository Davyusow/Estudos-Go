package main

import "fmt"

type cachorroquente int
var b cachorroquente //semelhante ao alias do bash e typedef de C, mas pode ser usado de várias formas

func main() {
    fmt.Printf("%T b\n",b)

    //mas ela não é um int, ao tentar comparar por exemplo da pau
    num := 2
    //comparar := num == b    //da pau, tipos diferentes
    // O que devemos fazer é um conversion, semelhante ao cast de Java mas aqui é diferente
    comparar := num == int(b)
    fmt.Println(comparar)
    fmt.Printf("%T b\n",b)

}
