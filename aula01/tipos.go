package main

//Os tipos em Go são estáticos
var valor uint8 = 1 //o bit de sinal conta não conta mais como sinal
var num int8 //ao declarar uma variável desta maneira, ela só poderar ter valor definido
// dentro de um code block
// Caso o tipo não tenha sido declarado, será definido dinâmicamente

var zero int  //ao não ser inicializada, pode sofrer diversas mudanças
//no caso do int é o 0 mesmo
// Para functions, pointers, e outras coisas que usam o endereço, o valor zero é nil: NULO

func main() {
    num = 16 //inicializa a variável e atribui
    println(num)
    num = 21 //atribui
    println(num)
    println(zero)

}
