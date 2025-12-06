package main

import "fmt"

func main(){
    arranjo := [5]int8{1,2,3,4,5} //tamanho fixo,
    fatia := []int{2,1,3,4}

    fmt.Println(arranjo)
    fmt.Println(fatia)
    fatia = append(fatia, 9)
    fmt.Println(fatia)

}
