package main

import "fmt"

func main () {
  var a [5]int = [5]int{32, 29, 78, 16, 81}
  var b = [5]int{32, 29, 78, 16, 81}
  c := [5]int{32,29,78,16,81}

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
