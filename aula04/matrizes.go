package main

import "fmt"

func main() {

    matriz := [][]int{
        []int{1,2,3},
        []int{4,5,6},
        []int{7,8,9},
    }

    fmt.Println(matriz) //matriz inteira
    fmt.Println(matriz[1])  //uma linha
    fmt.Println(matriz[2][2])   //um elemento

}
