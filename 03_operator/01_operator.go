package main

import "fmt"

func main () {
  var num1 uint8 = 3
  var num2 uint8 = 2

  fmt.Println(num1 + num2) // 5
  fmt.Println(num1 - num2) // 1
  fmt.Println(num1 * num2) // 6
  fmt.Println(num1 / num2) // 1
  fmt.Println(num1 % num2) // 1
  fmt.Println(num1 << num2) // 11(2) << 2 = 1110 = 6 + 4 + 2 + 0 = 12
  fmt.Println(num1 >> num2) // 11(2) >> 2 = 00.11 = 00 = 0
  fmt.Println(^num1)
  // ^00000011 = 11111100 = 128 + 64 + 32 + 16 + 8 + 4 + 0 + 0
  // 252
}
