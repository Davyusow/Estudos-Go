package main

import "fmt"

func main() {
    pedaco := make([]int,5,10) //tipo do slice, tamanho, capacidade máxima do slice

    pedaco[0],pedaco[1],pedaco[2],pedaco[3],pedaco[4] = 1,2,3,4, 5

    fmt.Println(pedaco, len(pedaco), cap(pedaco))

    pedaco = append(pedaco, 10)
    pedaco = append(pedaco, 10)
    pedaco = append(pedaco, 10)
    pedaco = append(pedaco, 10)
    pedaco = append(pedaco, 10)
    pedaco = append(pedaco, 10)

    fmt.Println(pedaco, len(pedaco), cap(pedaco))   //após estourar a capacidade, dobrou ela


}
