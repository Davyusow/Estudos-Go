package main

import "fmt"

func main() {
    var arranjo [5]int

    arranjo[0] = 10
    arranjo[1] = 11
    arranjo[2] = 12
    arranjo[3] = 13
    arranjo[4] = 14

    for indice, valor := range(arranjo) {
        fmt.Printf("Array[%d]: %2d\n",indice,valor)
    }

    fmt.Printf("\nTipo: %T\n",arranjo)
}
