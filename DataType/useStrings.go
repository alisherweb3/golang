package main

import "fmt" s "strings" "unicode"

var f = fmt.Printf



func main() {
  upper := s.ToUpper("Hello there!")
  f("To Upper: %s\n", upper)
  f("To Lower: %s\n", s.ToLower("Hello THERE"))
  f("%s\n", s.Title("tHis wiLL be a title!"))
  f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
  f("EqualFold: %v\n", sEqualFold("Mihalis", "Mihali"))
  
  
  
  
  f("Prefix: %v\n", s.HasPrefix(""))
  f("Prefix: %v\n", s.HasPrefix(""))
  
                                F("Prefix:
