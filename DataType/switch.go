package main

import (
  "fmt"
  "os"
  "regexp"
  "strconv"
)

func main() {
  arguments := os.Args
  if len(arguments) < 2 {
    fmt.Println("usage: switch number")
    os.Exit(1)
  }
  
  
  
  number, err := strconv.Atoi(arguments[1])
  if err != nil {
    fmt.Println("This value is not  an integer:", number)
  } else {
    switch {
    case number < 0:
      fmt.Println("Less than Zero")
    case number > 0:
      fmt.Println("Bugger than Zero!")
    default:
      fmt.Println("Zero!)
                  }
                  
