package main

var x interface{}

func main(){
    x = true

    switch x.(type) {
        case bool:
            println("Booleano")
        case int:
            println("Int")
        case string:
            println("String")
        default:
            println("<Nil>")
    }
}
