package main

import "fmt"

func f() {
  a := [...]int{1, 2}

  for i := 0; i <5; i++ {
    fmt.Println(a[i])
  }
}

func main() {
  f()

  fmt.Println("Hello, world")
}
