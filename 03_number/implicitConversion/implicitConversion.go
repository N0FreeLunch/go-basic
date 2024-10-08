package main

import "fmt"

func main () {
  var num1 int = 3
  var num2 float32 = 2.2

  // fmt.Println(num1 + num2) // error
  fmt.Println(float32(num1) + num2)
  fmt.Println(num1 + int(num2))

  var num3 uint16 = 10
  var num4 uint32 = 80000

  fmt.Println(num3 + uint16(num4))
  // fmt.Println(num3 + num4) // error
}
