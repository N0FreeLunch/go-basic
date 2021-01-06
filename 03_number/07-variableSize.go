package main

import "fmt"
import "unsafe"

func main() {
  var num1 int8 = 1
  var num2 int16 = 1
  var num3 int32 = 1
  var num4 int64 = 1

  fmt.Print(unsafe.Sizeof(num1))
  fmt.Println("byte")
  fmt.Print(unsafe.Sizeof(num2))
  fmt.Println("byte")
  fmt.Print(unsafe.Sizeof(num3))
  fmt.Println("byte")
  fmt.Print(unsafe.Sizeof(num4))
  fmt.Println("byte")
}
