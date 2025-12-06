package main

import "fmt"

const (
    HOJE = 2025
    AMANHA1 = HOJE + iota
    AMANHA2 = HOJE + iota
    AMANHA3 = HOJE + iota
    AMANHA4 = HOJE + iota
)

func main(){

    println(HOJE)
    println(AMANHA1,AMANHA2,AMANHA3,AMANHA4)
    fmt.Printf("%#x\n",31337)
}
