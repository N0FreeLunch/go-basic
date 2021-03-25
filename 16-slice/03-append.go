package main

import "fmt"

func main () {
  a := []int{1, 2, 3}
  b := []int{4, 5, 6}

  a = append(a, 4, 5, 6)
  b = append(a, 4, 5, 6)

  fmt.Println(a)
  fmt.Println(b)
}
