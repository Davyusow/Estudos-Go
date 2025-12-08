package main

import "fmt"

func main() {
    //é possível ter canais que só enviam, e canais que só devolvel informações (send & receive)

    canal := make(chan int) //declara como receive

    go send(canal) //neste caso o canal se molda a depender do parâmetro da função
    receive(canal)

}


func send(s chan<- int) {
    s <- 42
}

func receive(r <-chan int) {
    fmt.Println("O valor é:",<-r)
}
