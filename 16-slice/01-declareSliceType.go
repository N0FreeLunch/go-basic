package main

import "fmt"

func main () {
  var a []int = make([]int, 5)
  var b = make([]int, 5)
  c := make([]int, 5)

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
