package main
import "fmt"


func repeat(w string, x int) {
    for i := 0; i < x; i++ {
        fmt.Println(w)
    }
}

func main() {
    var w string
    fmt.Scanln(&w)
    var x int
    fmt.Scanln(&x)

    repeat(w, x)

}
