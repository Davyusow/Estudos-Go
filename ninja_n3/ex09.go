package main

const (
    vol = "vôlei"
    fut = "futebol"
    f1 = "fórmula 1"
)

func main(){
    x := "fórmula 1"

    switch (x){
        case "fórmula 1":
            println(f1)
        case "futebol":
            println(fut)
        case "vôlei":
            println(vol)
        default:
            println("Nenhum desses vai entender")
    }
}
