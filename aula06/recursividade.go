package main

func fatorial(x int) int {  // recursividade toma mais memória
    if (x == 1){
        return  x
    }
    return x * fatorial(x-1)
}

func main() {
    println(fatorial(5))

}
