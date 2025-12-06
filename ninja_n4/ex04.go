package main

import "fmt"

func main() {
    lista := []int{42,43,44,45,46,47,48,49,50,51}
    fmt.Println(lista)

    lista = append(lista, 52)
    fmt.Println(lista)

    lista = append(lista, 53,54,55)
    fmt.Println(lista)

    y := []int{56,57,58,59,60}

    lista = append(lista, y...)
    fmt.Println(lista)

}
