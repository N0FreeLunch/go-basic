package main

import "fmt"

func main () {
   a := [5]int{1, 2, 3, 4, 5}
   b := a
   a[2] = 0
  fmt.Println(a)
  fmt.Println(b)
}
