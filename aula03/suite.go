package main

const (
    seg = "Segunda"
    ter = "Terca"
    qua = "Quarta"
    qui = "Domingo"
    sex = "Sexta"
    sab = "Sábado"
    dom = "Domingo"
)

func main() {

    dia := 7

    switch (dia) { // diferente de outras linguagens, go não tem fallthrough por padrão(ainda bem)
        case 1 :
            println(dom)
        case 2:
            println(seg)
        case 3:
            println(ter)
        case 4:
            println(qua)
        case 5:
            println(qui)
        case 6:
            println(sex)
            fallthrough //aqui faz rolar o próximo caso também
        case 7:
            println(sab)
            fallthrough
        default:
            println("Desconhecidooo")
    }
}
