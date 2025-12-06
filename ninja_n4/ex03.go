package main

import "fmt"

func main() {
    valores := []int{10,20,30,40,50,60,70,80,90,100}

    for i, v := range valores {
        fmt.Println(i,v)
    }

    x := valores[0:3]
    y := valores[4:]
    z := valores[2:7]
    a := valores[2:len(valores)-1]

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    fmt.Println(a)

}
