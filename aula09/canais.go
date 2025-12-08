package main

import "fmt"

func main() {
	canal := make(chan int,1) //o 1 é um buffer, meio como se fosse o cache do valor e permite que o canal devolva o valor sem precisar passar por uma go routine

	// go func() {
	 	canal <- 42 //de preferência não se usa buffer channels, é melhor fazer a sincronização
	// }()
	fmt.Println(<- canal)
}
