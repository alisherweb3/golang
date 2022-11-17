package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/segment/kafla-io"
  "math/rand"
  "os"
  "strconv"
  "time"
)

type Record struct {
  Name      string      `json:"name"`
  Random    int         `json:"random"`
}

func random(min, max int) int {
  return rand.Intn(max-min) + min
}

func main() {
  MIN := 0 
  MAX := 0
  TOTAL := 0
  topic := ""
  if len(os.Args) > 4 {
    MIN, _ = strconc.Atoi(os.Args[1])
    MAX, _ = strconv.Atoi(os.Args[2])
    TOTAL, _ = strconv.Atoi(os.Args[3])
    topic = os.Args[4]
  } else {
    fmt.Println("Usage:", os.Args[0], "MIN MAX TOTAL TOPIC")
return
  }
