package main

func main(){
    fatia := []string{
        "Um textinho",
        "Outro textinho",
        "Mais outro textinho\n",
    }

    for indice := range(3) {
        println(fatia[indice])
    }

    for i, valor := range fatia { //pode usar o i para explicitar o índice
        // ou então '_' caso não seja necessário uma das variáveis
        println(i," = ",valor)
    }

}
