package main

import "fmt"

const (
  a = 1 << iota
  b = 1 << iota
  c = 1 << iota
  d = 1 << iota
)

const (
  aa = iota * 30
  bb = iota * 30
  cc = iota * 30
  dd = iota * 30
)

func main () {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(aa)
  fmt.Println(bb)
  fmt.Println(cc)
  fmt.Println(dd)
}
