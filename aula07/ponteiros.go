package main

import "fmt"

func main() {
    x := 10

    y := &x //recebe o endereço igualzinho em C

    fmt.Println(y)
    fmt.Println(*y)
    fmt.Printf("%T\n",x)
    fmt.Printf("%T\n",y)
    fmt.Printf("%T\n",*y)
    *y++    //também altera o valor original
    fmt.Printf("%v\n",x)
    fmt.Printf("%v\n",y)
    fmt.Printf("%v\n",*y)

    fmt.Println("Valor de X:",x)
    mult2(&x)    //muito mais eficiente
    fmt.Println("Valor de X:",x)
}


func mult2(x *int) {
    *x *= 2
}
