package main

import "fmt"

type veiculo struct {
    portas int
    cor string
}

type caminhonete struct {
    veiculo
    tracao_4x4 bool
}

type sedan struct {
    veiculo
    modeloLuxo bool
}

func main(){

    hilux := caminhonete {
        veiculo {
            portas: 12,
            cor: "Branco",
        },
        true,
    }

    voyage := sedan {
        veiculo {
            4,
            "Cinza",
        },
        false,
    }

    fmt.Println(hilux)
    fmt.Println(voyage)



}
