package main

func main() {
    defer println("Fim!")   //roda por último

    defer println("Penúltimo") //a ordem é, quem foi definido primeiro vem por último
    for i := range(11) {
        for j := range(11){
            println(i,"x",j,"=",i*j)
        }
    }
}
