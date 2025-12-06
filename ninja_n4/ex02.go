package main

import "fmt"

func main() {
    valores := []int{10,20,30,40,50,60,70,80,90,100}

    for i, v := range valores {
        fmt.Println(i,v)
    }
}
