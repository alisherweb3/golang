package main

import "fmt"

func main() {
    //your code goes here
    var t float32
    fmt.Scanln(&t)
    if t > 99.5 {
        fmt.Println("Fever")
    } else {
        fmt.Println("Allowed")
    }
}
