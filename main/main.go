package main

import (
	"encoding/json"
	"fmt"
	"strdiff"
)

func main() {
  a := "intention"
  b := "execution"
  result := strdiff.Diff(a, b)
  data, err := json.Marshal(result)
  if err != nil {
    fmt.Printf("%+v\n", err)
  }
  fmt.Println(string(data))
}
