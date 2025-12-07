package main

import (
	"fmt"
)

func main() {

	x := 12

	y := func(x int) int { //sim, tem como uma variável receber uma função💀
		return x * 1000
	}
	fmt.Println(x, "vezes 1000 é:", y(x))

	z := doido(x)
	a := z(10)  //isso é mt feio véi

	c := doido(2)(3)
	println(c)

	println(a)


}

func doido(x int) func(int) int { // a função mais feia que escrevi na vida, mas é legal saber que da pra fazer isso

    return func(i int) int {
        y := i * 10
        return  y
    }
}
