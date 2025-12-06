package main

import "fmt"

func main() {
    x := 10

    fmt.Printf("%d\n%#b\n%#x\n",x,x,x)

    y := x << 1

    fmt.Printf("%d\n%#b\n%#x\n",y,y,y)
}
