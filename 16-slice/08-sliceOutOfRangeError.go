package main

import "fmt"

func main () {
  var a = make([]int, 5, 10)

  fmt.Println(a[4])
  fmt.Println(a[5])
  fmt.Println(a[8])
}
