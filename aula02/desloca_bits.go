package main

import "fmt"

func main() {
	x := 0b00000011

	fmt.Printf("%3d %#12b \n", x<<2, x<<2)
	fmt.Printf("%3d %#12b \n", x<<4, x<<4)
	fmt.Printf("%3d %#12b \n", x<<8, x<<8)
	fmt.Printf("%3d %#12b \n", x>>1, x>>1)
}
