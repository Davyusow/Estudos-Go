package main

func main() {

    x := 10
    for {
        if x == 20{
            break
        } else {
            println(x)
            x++
            //continue
        }
    }
    println("Sai do loop")
    x = 0
    for x < 20{
        if (x % 2 == 1) {
            println(x)
        }
        x++
        continue //pula diretamente pro próximo loop(i++)
    }
}
