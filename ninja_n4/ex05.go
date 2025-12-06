package main

import "fmt"

func main() {
    lista := []int{42,43,44,45,46,47,48,49,50,51}

    y := append(lista[:3], lista[6:]...)

    fmt.Println(y)
}
