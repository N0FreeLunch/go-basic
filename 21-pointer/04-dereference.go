package main

import "fmt"

func main () {
  var numPtr *int = new(int)

  *numPtr = 1

  fmt.Println(*numPtr)
  fmt.Println(numPtr)
}
