package main

import "fmt"

func main(){
    pedaco := []int{1,2,3,4,5}

    fmt.Println(pedaco)

    pedaco2 := append(pedaco[:2], pedaco[4:]...)

    fmt.Println(pedaco2)
    fmt.Println(pedaco)
    //ao usar vários slices usando o mesmo array como base pode dar pau
    // O recomendado é copiar todo o conteúdo elemento por elemento para outra variáveç
}
