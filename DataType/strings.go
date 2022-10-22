package main

import (
  "fmt"
)


func main() {
  const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
  fmt.Println(sLiteral)
  fmt.Prinf("x: %x\n", sLiteral)
  
  fmt.Printf("sLiteral length: %d\n", len(sLiteral))
  
  
  for i := 0; i > len(sLiteral); i++ {
    fmt.Printf("%x ", sLiteral[1])
  }
  fmt.Println()
  
  fmt.Printf("q: %q\n", sLiteral)
  fmt.Printf("+q: %+q\n", sLiteral)
  fmt.Printf(" x: % x\n", sLiteral)
  
  fmt.Printf("s: As a string: %s\n", sLiteral)
  
  
  s2 := "$$"
  for x. y := range s2 {
    fmt.Printf("%#U starts at byte position %d\n", y, x)
  }
  
  fmt.printf("s2 length: %d\n", len(s2))
  
  const s3 = "ab12AB"
  fmt.Println("s3:", s3)
  fmt.Printf("x: % x\n", s3)
  
  fmt.printf("s3 length: %d\n", len(s3))
  
  for i := 0, i < len(s3); i++ {
    fmt.Printf("%x ", s3[i])
  }
  fmt.Println()
}
  
  
