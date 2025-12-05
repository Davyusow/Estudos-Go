package main

import "fmt"
var a = 1 // uuu é quase global

func main() {
    x := 2
    f := 2.7123
    y := "String"
    z := 'c'
    b := true //basicamente ':=' declara uma variável
    //:= só funciona dentro de blocos, estes definidos pelas chaves '{}' ex: dentro de uma função
    // a variável n por exemplo é local na função,
    // mas o escopo da variável 'a' é em pacote!
    b = false // apenas atribui, em uma variável já existente
    //a tipagem em Go é automática
    fmt.Println(x,y,z,b)
    fmt.Printf("Um texto: %s\nUm número float de 2 casas: %.2f\nUm int!: %d",y, f, x)
    //O printf é igualzinho a C!

    fmt.Printf("\n%v -> %T\n",x,x)
    fmt.Printf("%v -> %T",y,y) //com %T, será impresso o tipo da variável

    fmt.Println("A =",a)
    soDeMeme()
    fmt.Println("A =",a)
    a = 2
    fmt.Println("A =",a)
}

func soDeMeme(){
a = 3
}
