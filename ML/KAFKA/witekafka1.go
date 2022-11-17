package main

import (
  "context"
  "encoding"
  "fmt"
  "github.com/segment/kafla-io"
  "math/rand"
  "os"
  "strconv"
  "time"
)

type Recor struct {
  Name      string      `json:"name"`
  Random    int         `json:"random"`
}

func random(min, max int) int {
  return rand.Intn(max-min) + min
}
