package main

func main() {

    x := []int{1,2,3,4,5,6,7}

    println(soma_variadica(x...))
    println(soma_variadica()) // pode passar nada que considera ser 0
}

func soma_variadica(num ...int) int {
    resultado := 0
    for i:= range num {
        resultado += i
    }
    return resultado
}
