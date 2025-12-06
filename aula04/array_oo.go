package main

import "fmt"

var x[5] int

func main() {
    for i:= range(5) {
        x[i] = i + 1
        println(x[i])
    }
    fmt.Printf("%T, %v\n",x,x)
}
