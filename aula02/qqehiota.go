package main

const (
    x = iota
    y = iota
    _ = iota
    _ = iota //ignora o 2, e 3
    z = iota //criou valores sucessivos, inteiros não tipados
) //serve para quando você quer as contantes diferente das outras

//honestamente achei meio dispensável mas deve haver alguma coisa muito útil com isso

func main(){
    println(x,y,z)
}
