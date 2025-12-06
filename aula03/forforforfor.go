package main

func main() {

    i := 10

    for i = 0;i < 10; i++ {
        println(i)
    }
    println(i)


    x := 0

    for x <= 10 { //funciona como o while, já que em go só existe o for
        println("X é menor do que 10!")
        x++
    }
    println("X é maior do que 10!")
    println(x)

    for { //go tem uma proteção contra processos infinitos (process took too long)

    }
}
