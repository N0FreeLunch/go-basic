package main

import "fmt"
import "math"

func main() {
  var num1 uint8 = math.MaxUint8 + 1
  var num2 uint16 = math.MaxUint16 + 1
  var num3 uint32 = math.MaxUint32 + 1
  var num4 uint64 = math.MaxUint64 + 1

  fmt.Println(num1)
  fmt.Println(num2)
  fmt.Println(num3)
  fmt.Println(num4)
}
