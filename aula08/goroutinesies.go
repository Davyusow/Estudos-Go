package main

import (
	"fmt"
	"runtime"
	"sync"
)

    var wg sync.WaitGroup = sync.WaitGroup{}
func main() {

    fmt.Println(runtime.NumCPU())
    wg.Add(2)   //importantissimo definir quantas go routines existem no arquivo
    defer wg.Wait() //ótimo momento pra usar isso

    go func1() //cria uma goroutine, semelhante a uma thread
    // dessa forma, o programa acaba antes mesmo da função rodar
    // porém leva um tempo para o kernel alocar a thread, então precisa das duas funções serem uma go routine para ter concorrência
    fmt.Println("GO routines:",runtime.NumGoroutine())
    go func2()
    //infelizmente(na verdade infelizmente) é tão rápido que não da pra perceber bem
}

func func1() {
    for i := range 10 {
        fmt.Println("func1:",i+1)
    }
    wg.Done()
}

func func2() {
    for i := range 10 {
        fmt.Println("func2:",i+1)
    }
    wg.Done()
}
