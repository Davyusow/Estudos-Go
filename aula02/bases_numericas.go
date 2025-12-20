package main

import "fmt"

func main() {
	i := 0b1000
	a := 0x8
	c := 0o10

	fmt.Printf("%d, %#b, %#x\n", a, a, a)
	fmt.Printf("%d, %#b, %#x\n", i, i, i)
	fmt.Printf("%d, %#b, %#o\n", c, c, c)
}
