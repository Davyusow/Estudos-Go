package main

import "fmt"

func par(num int) int {
    return 2*num
}

func par2(num int) (int, string) {
    return 2*num, fmt.Sprintln("é o",num,"º número par")
}

func main() {

    num := par(10)

    val, txt := par2(10)

    println(num)
    println(val,txt)

}
