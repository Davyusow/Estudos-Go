package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

    var contador int64 = 0
    total_goroutines := 15

    wg.Add(total_goroutines)


    for i := 0 ; i < total_goroutines; i++ {
        go func(){
            defer wg.Done()

            atomic.AddInt64(&contador, 1)
            runtime.Gosched()
            fmt.Println("Contador: ",atomic.LoadInt64(&contador))

        }()
    }

    wg.Wait()   //lembrar disso! só usar normalmente a variável quando o waitSync tiver feito tudo
    fmt.Println(contador)

}
