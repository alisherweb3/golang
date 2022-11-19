package main

import (
  "flag"
  "fmt"
  "sync"
)

func main() {
  n := flag.Int("n", 20, "Number of goroutines")
  flag.Parse()
  count := *n
  fmt.Printf("Going to create %d goroutines.\n", count)
  
  var waitGroup sync.WaitGroup
  
  type WaitGroup struct {
    noCopy
    state1    [12]byte
    sema      uint32
  }
  
  fmt.Printf("%#v\n", waitgroup)
  for i := 0; i< count; i++ {
    waitGroup.Add(1)
    go func(x int) {
      defer waitGroup.Done()
      fmt.Printf("%d ", x)
    }(i)
  }
  fmt.Printf("%#v\n", waitGroup)
  waitGroup.Wait()
  fmt.Println("\nExiting...")
}
