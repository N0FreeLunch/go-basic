package main

import "fmt"

func main () {
  var numPtr *int = new(int)

  numPtr++
  numPtr = 0xc0820062d0

  fmt.Println(numPtr)
}
