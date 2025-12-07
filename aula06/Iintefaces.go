package main

import "fmt"

type pessoa struct {
    nome string
    idade int
}

type dentista struct {
    pessoa
    dentes_arrancados int
    salario float64
}

type arquiteto struct {
    pessoa
    tipo_de_construcao string
    salario float64
}

func (p pessoa) imprimir() {
    fmt.Println("Nome: ",p.nome)
    fmt.Println("Idade: ",p.idade)
}

func (p dentista) imprimir() {
    fmt.Println("Nome:",p.nome)
    fmt.Println("Idade:",p.idade)
    fmt.Println("Dentes Arrancados:",p.dentes_arrancados)
    fmt.Println("Salário:",p.salario)
}

func (p arquiteto) imprimir() {
    fmt.Println("Nome:",p.nome)
    fmt.Println("Idade:",p.idade)
    fmt.Println("Tipo de Construção:",p.tipo_de_construcao)
    fmt.Println("Salário:",p.salario)
}

type gente interface {  //implicitamente os tipos que tiverem estes métodos serão da interface gente
    imprimir()
}

func serHumano(g gente) {   //portanto poderão ser chamados como gente aqui dentro
    //g.imprimir()    //é possível fazer um switch para cada tipo aqui dentro
    switch g.(type) {
        case pessoa:
            g.imprimir()
            println("É uma pessoa que implementa gente")
        case arquiteto:
            g.imprimir()
            println("É um arquiteto que implementa gente")
        case dentista:
            g.imprimir()
            println("É um dentista que implementa gente")
        default:
            println("Tipo inválido!")
    }
}

func main() {

    arq1 := arquiteto{
        pessoa{
            "Fulano",
            21,
        },
        "Cívil",
        5000.00,
    }

    dent1 := dentista{
        pessoa{
            "Fulano",
            21,
        },
        500,
        4500.00,
    }

    arq1.imprimir()
    println()
    dent1.imprimir()

    println()
    serHumano(arq1)
    println()
    serHumano(dent1)
}
