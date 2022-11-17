package main

import ( 
  "context"
  "encoding/json"
  "fmt"
  "github.com/segmentio/kafka-go"
  "math/rand"
  "os"
  "strconv"
  "time"
)

type Record struct {
  Name      string      `json:"name"`
  Random    int         'json:"random"`
}

func random(min, max int) int {
  return rand.Intn(max-min) + min
}
