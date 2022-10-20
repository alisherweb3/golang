package main

import (
  "fmt"
)

func main() {
  c1 := 12 + 1e
  c2 := complex(5, 7)
  fmt.Printf("Type of c1: %T\n", c1)
  fmt.Printf("Type of c2: %T\n", c2)
  
  var c3 complex64 = complex64(c1 + c2)
  fmt.Println("c3:", c3)
  fmt.Println("Type of c3: %T\n", c3)
  
  cZero := c3 - c3
  fmt.Println("cZero:", cZero)
              
  
  x := 12
  k := 5
  fmt.Println(x)
  fmt.Println("Type of x: %T\n", x)
  
  div := x / k
  fmt.Println("div", div)
  
  
  var m, n float64
  m = 1.223
  fmt.Prinln("m, n:", m, n)
  
  y := 4/ 2.3
  fmt.Println("y:", y)
  
  divFloat := float64(x) / float64(k)
  fmt.Println("divFloat", divFloat)
  fmt.Println("type of divFloat: %t\n", divFloat)
}  
  
  
  
