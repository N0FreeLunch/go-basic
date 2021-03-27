package main

import "fmt"

func SumAndDiff (a int, b int) (sum int, diff int) {
  sum = a + b
  diff = a - b
  return
}

func main () {
  sum1, diff1 := SumAndDiff(6, 2)
  fmt.Println(sum, diff)
}
