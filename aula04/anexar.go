package main

import "fmt"

func main() {
    pedaco := []int{1,2,3,4}
    outro_pedaco := []int{9,10,11,12}

    fmt.Println(pedaco)

    pedaco = append(pedaco, 5,6,7,9) //append anexa os elementos em uma slice no fim de uma slice

    fmt.Println(pedaco)

    pedaco = append(pedaco, outro_pedaco...) //os três pontinhos quer dizer pra pegar o conteúdo da fatia
    //unfurl, desenrolar, enumeration

    fmt.Println(pedaco)
}
