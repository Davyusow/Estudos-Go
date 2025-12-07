package main

func main() {
    println(soma(10,22))
    basica()
    x, y :=  soma_maior_que4(10,22)
    println(x)
    println(y)
    println(soma_variadica(1,2,3,4,5,6,7,8,9,10))
}
                        //esse int simboliza o tipo do retorno
func soma(v1 int, v2 int) int { //ou "v1 , v2 int"
    return v1 + v2
}

func basica() { //quando não tem retorno pode ser omitido
    println("texto")
}

func soma_maior_que4(v1 int, v2 int) (int, bool) { //ou "v1 , v2 int"
    somas := v1 + v2
    maior :=  somas > 4
    return somas, maior
}

func soma_variadica(num ...int) int { //recebe vários valores
    resultado := 0
    for i:= range num {
        resultado += i
    }
    return resultado
}
