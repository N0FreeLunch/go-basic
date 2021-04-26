package main

import "fmt"

var t interface{}

type Person struct {
  name string
  age int
}

func main () {
  t = Person{"Maria", 20}
  if v, ok := t.(Person); ok {
    fmt.Println(v, ok)
  }
}
