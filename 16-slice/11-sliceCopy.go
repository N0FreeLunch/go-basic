package main

import "fmt"

func main () {
  a := []int{1,2,3,4,5}
  b := make([]int, 3)

  fmt.Println(b)

  copy(b,a)
  b[0] = 9

  fmt.Println(a)
  fmt.Println(b)

}
