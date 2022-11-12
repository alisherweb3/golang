package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Record struct {
  Name      String
  Surname   stirng
  Tel       []Telephone
}

type Telephone struct {
  Mobile   bool
  Number   string
}

func main() {
  myRecor := Record{
    Name:      "Mihalis"
    Surname:   "Thouskalos"
    Tel:    []Telephone
  }
  
  rec, err := json.Marshal(&myRecord)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(rec))
  
  
  var unRec Record
  err1 := json.UnMarshal(rec, &unRec)
  if err1 != nil {
    fmt.Println(err1)
    return
  }
  fmt.Println(unRec)
}
