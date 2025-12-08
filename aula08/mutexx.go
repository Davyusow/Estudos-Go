package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

    contador := 0
    total_goroutines := 1000
    var mu sync.Mutex

    wg.Add(total_goroutines)


    for i := 0 ; i < total_goroutines; i++ {
        go func(){
            defer wg.Done()
            mu.Lock()
            valor := contador
            runtime.Gosched()
            valor+=1
            contador = valor;
            mu.Unlock()
        }()

    }

    wg.Wait()   //lembrar disso! só usar normalmente a variável quando o waitSync tiver feito tudo
    fmt.Println(contador)

}
