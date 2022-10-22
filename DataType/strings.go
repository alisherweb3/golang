package main

import (
  "fmt"
)


func main() {
  const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
  fmt.Println(sLiteral)
  fmt.Prinf("x: %x\n", sLiteral)
  
  fmt.Printf("sLiteral length: %d\n", len(sLiteral))
  
  
