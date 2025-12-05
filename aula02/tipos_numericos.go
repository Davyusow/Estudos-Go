package main

import (
	"fmt"
	"runtime"
)

var x uint8 //tem 8, 16, 32 e 64 bits, lembra rust
var z int64 //tipo long de C
var y float32  // tem 32 e 64 bits (normalemente é melhor usar o 64 para maior precisão)

func main () {
    a := '2'
    b := 'b'
    c := '%' //um rune é subjacente de um int32

    println("a '2': ",int(a)) //imprime o valor do caractere e não '2'
    println("b 'b': ",int(b))
    println("c '%': ",int(c))


    fmt.Println(runtime.GOARCH) //imprime a arquitetura do computador
    fmt.Println(runtime.GOOS) //imprime o sistema Operacional

    x = 255
    //como x é um uint8, ele tem escopo de 0 a 255, ao passar disso é feito um overflow, ele reseta o escopo voltando à 0
    //mas go não permite overflow definido diretamente, mas em tempo de execução é possível mostrar isso
    println(x) //255
    println(x+1) //256
}
