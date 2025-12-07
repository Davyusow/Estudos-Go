package main

import (
	"fmt"
)


func soma2(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func pares(f func(x ...int) int, y ...int) int { //callback é esse troço, de passar uma função como argumento
    //pode ser útil em situações que a depender de um gerador, por exemplo: dificuldades
    // pode alterar diretamente no funcionamento da função
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func impares(f func(x ...int) int, y ...int) int {
    var pedaco []int
    for _, valor := range y {
        if valor % 2 == 1 {
            pedaco = append(pedaco, valor)
        }
    }
    total := f(pedaco...)
    return total
}

func main() {
	t := pares(soma2, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)

	total := impares(soma2, []int{1,2,3,4,5,6,7,9,10}...)

	fmt.Println(t)
	fmt.Println(total)
}
