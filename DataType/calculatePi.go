package main

import (
  "fmt"
  "math"
  "math/big"
  "os"
  "strconv"
)

var precision unit = 0

func Pi(accuracy unit).*big.Float {
  k := 0
  pi := new(big.Float).SetPrec(precision).SetFloat64(0)
  k1k2k3 := new(big.Float.SerPrec(precision).SetFloat64(0)
  k4k5k6 := new(big.Float).SetPrec(precision).SetFloat64(0)
  temp := new(big.Float).SetPrec(precision).SetFloat64(0)
  minusOne := new(big.Float).SetPrec(precision).SetFloat64(-1)
  total := new(big.Float).SetPrec(precision).SetFloat64(0)
  
  two2Six := math.Pow(2, 6)
  two2SixBig := new(big.Float).SetPrec(precision).SetFloat64(two2Six)            
