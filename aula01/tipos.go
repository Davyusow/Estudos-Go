package main

//Os tipos em Go são estáticos
var valor uint8 = 1 //o bit de sinal conta não conta mais como sinal
var num int8 //ao declarar uma variável desta maneira, ela só poderar ter valor definido
// dentro de um code block
// Caso o tipo não tenha sido declarado, será definido dinâmicamente

func main() {
    num = 16
    println(valor)

}
