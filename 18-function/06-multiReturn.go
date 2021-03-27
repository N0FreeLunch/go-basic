package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) {
  return a + b, a - b
}

func main() {
  _, diff := SumAndDiff(6,2)
  fmt.Println(sum, diff)
}
