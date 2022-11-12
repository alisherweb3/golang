packahe main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Record struct {
  Name string
  Surname string
  Tel []Telephone
}

type Telephone struct {
  Mobile bool
  Number string
}

func saveToJSON(filename *os.File, key interface{}) {
  encodeJson := json.NewEncoder(filename)
  err := encodeJSON.Encode(key)
  if err != nil {
    fmt.Println(err)
    return
  }
}

func main() {
  myRecord := Record{
    name:          "Mihalis"
    Surname:       "Thouskalos"
    Tel: []Telephone{Telephone{Mobile: True, Number: "1234-567"},
                     Telephone{Mobile: true, Number: 1234-abcd"},
                               Telephone{Mobile: true, Number: abcc-567"}
                                        },
                              },
                    },
  }
  saveToJSON(os.Stdout, myRecord)
  
    
