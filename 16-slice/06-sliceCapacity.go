package main

import "fmt"

func main () {
  var a = make([]int, 5, 10)
  b := append(a, 1,2,3,4,5,6,7)
  c := append(b, 8,9,10)
  d := append(c, 11,12,13,14,15)
  e := append(d, 16)

  fmt.Println(a)
  fmt.Println(len(a))
  fmt.Println(cap(a))
  fmt.Println(b)
  fmt.Println(len(b))
  fmt.Println(cap(b))
  fmt.Println(c)
  fmt.Println(len(c))
  fmt.Println(cap(c))
  fmt.Println(d)
  fmt.Println(len(d))
  fmt.Println(cap(d))
  fmt.Println(e)
  fmt.Println(len(e))
  fmt.Println(cap(e))
}
