package main

import "fmt"

func main () {
  var numPtr *int = new(int)
  fmt.Println(numPtr)
}
