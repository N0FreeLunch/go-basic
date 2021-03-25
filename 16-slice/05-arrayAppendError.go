package main

import "fmt"

func main () {
  var a [5]int = [5]int{32, 29, 78, 16, 81}
  a = append(a, 100)

  var b = [5]int{32, 29, 78, 16, 81}
  b = append(b, 100)

  c := [5]int{32,29,78,16,81}
  c = append(c, 100)

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
