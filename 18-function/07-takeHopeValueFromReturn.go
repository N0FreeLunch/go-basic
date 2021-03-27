package main

import "fmt"

func hello() (int, int, int, int, int) {
  return 1, 2, 3, 4, 5
}

func main() {
  a, _, c, _, e := hello()
  fmt.Println(a, c, e)
}
