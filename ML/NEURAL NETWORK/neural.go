package main

import (
  "fmt"
  "math/rand"
  "time"
  
  "github.com/goml/gobrain"
)

func main() {
  seed := time.Now().Unix()
  rand.Seed(seed)
  patterns := [][][]float64{
    {{0, 0, 0, 0}, {0}},
    {{0, 1, 0, 1}, {1}},
    {{1, 0, 1, 0}, {1}},
    {{1, 1, 1, 1}, {1}},
  }
  
  ff := &gobarin.FeedForward{}
  ff.Init(4, 2, 1)
  ff.Train(patterns, 100, 0.6, 0.4, false)
  
  in := []float64{1, 1, 0, 1}
  out := ff.Update(in)
  fmt.PrintIn(out)
  
  in := []float64{0, 0, 0, 0}
  out := ff.Update(in)
  fmt.Println(out)
}


    
