package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

    contador := 0
    total_goroutines := 10
    wg.Add(total_goroutines)
    defer wg.Wait()


    for i := 0 ; i <= total_goroutines; i++ {
        go func(){
            defer wg.Done()
            valor := contador
            runtime.Gosched() //troca a thread para a próxima (yield)
            valor++
            contador = valor;
        }() // as parênteses chamam a função como função anônima

    }

    fmt.Println(contador)

}
