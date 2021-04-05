package main

import "fmt"

type Rectangle struct {
  width int
  height int
}

func (_ Rectangle) dummy() {
  fmt.Println("dummy");
}

func main() {
  var  r Rectangle
  r.dummy()
}
