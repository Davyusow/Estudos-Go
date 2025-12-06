package main

import "fmt"

func main(){

    for hora := 0;hora < 24; hora++ {
        for minuto := 0;minuto < 60; minuto++ {
            fmt.Printf("São %02d:%02d\n",hora,minuto)
        }
    }
}
