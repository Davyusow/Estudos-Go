package main

import "fmt"

type numero int
var x numero

func main(){

    fmt.Printf("X: %v\n",x)
    fmt.Printf("X: %T\n",x)
    x = 42
    fmt.Printf("X: %v\n",x)


}
