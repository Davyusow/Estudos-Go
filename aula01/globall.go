package main

import "fmt"

var zaun = 20 //var só cria a variável em nível de pacote

func main() {
	y := 10
	numeros(y)
}

func numeros(x int) { //diferente de outras linguagens, o tipo vem depois
	fmt.Printf("Um texto de outra função!\nNúmero: %d\n", x)
	fmt.Println("Número global:", zaun) //como é uma variável de escopo do pacote, ela pode ser chamada em qualquer código go deste pacote!
}
