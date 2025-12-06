package main

import (
	"fmt"
)

func main(){
    cores := []string{"azul","amarelo","verde","vermelho","laranja"}

    fatia := cores[:len(cores)] //faz uma separação dos índices, lembra python

    fmt.Println(fatia)

    cores = append(cores[:2], cores[3:]... ) //apaga o verde cores[3]

    fmt.Println(cores)
}
