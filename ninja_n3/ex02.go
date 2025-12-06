package main

import "fmt"

var i rune = 65

func main() {

    for i <= 90{
        println(i)
        for j:=0 ; j < 3 ; j++{
            fmt.Printf("\t%#U \n",i) //caractere unicode
        }
        i++
    }
    println()
}
