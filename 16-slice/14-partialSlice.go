package main

import "fmt"

func main () {
  a := []int{1, 2, 3, 4, 5}
  b := a[0:5]

  fmt.Println(a)
  fmt.Println(b)

  fmt.Println(a[0:3])
  fmt.Println(a[1:3])
  fmt.Println(a[2:5])
}
