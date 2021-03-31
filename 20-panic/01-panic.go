package main

import "fmt"

func main () {
  a := [...]int{1, 2}

  for i := 0; i < 3; i++ {
    fmt.Println(a[i])
  }
}
