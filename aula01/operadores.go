package main

import "fmt"

func main(){
    a := 10
    b := 5
    soma := a + b
    subtracao := a - b
    multi := a * b
    divisao := a / b
    resto := a % b

    fmt.Println("Soma: ",soma,"\nSubtracao: ",subtracao,"\nMultiplicação:",multi)
    fmt.Println("Divisao",divisao,"\nResto ou MOD:",resto);

    v := true
    f := false

    e := v && f
    fmt.Println("operador 'e'",e);
    ou := v || f
    fmt.Println("ou ",ou);

    potencia := a^2
    fmt.Println("Potência: ",potencia);
    binario := 0b0010
    fmt.Println("Binários: ",binario);
    binario = binario << 2  //desloca 2 bits à esquerda
    fmt.Println("Binários: ",binario);
    binario = binario >> 2  //desloca 2 bits à direita
    fmt.Println("Binários: ",binario);

    booleano := 0b1010 &^ 0b0111 //um nand muito do feio, não entendi direito
    fmt.Printf("Booleano: %08b\n",booleano) //
    //tem a notação do += também, válida para TODAS as operações acima
    booleano &^= 0b1010 // olha que horror
    fmt.Printf("Booleano: %08b\n",booleano)

    a++ //só tem desse jeito e não invertido '++a' o que achei correto
    b--
    fmt.Println(a)
    fmt.Println(b)

    compara := e == ou
    fmt.Println("e == ou",compara) //operadores de comparação também igualzinho a C
    fmt.Printf("\n%d < %d: %d\n",a,b,a < b)
    fmt.Printf("\n%d > %d: %d\n",a,b,a > b)
    fmt.Printf("\n%d > %d: %d\n",a,b,!(a > b)) //pra mentir também é que nem C
    fmt.Printf("\n%d != %d: %d\n",a,b,a != b)
    fmt.Printf("\n%d <= %d: %d\n",a,b,a != b)
    fmt.Printf("\n%d >=  %d: %d\n",a,b,a != b)

}
