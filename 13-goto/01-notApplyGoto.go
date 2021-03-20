package main

import "fmt"

func main () {
  var a int = 1

  if a == 1 {
    fmt.Println("Error 1")
    return
  }

  if a == 2 {
    fmt.Println("Error 2")
    return
  }

  if a == 3 {
    fmt.Println("Error 1")
    return
  }

  fmt.Println(a)
  fmt.Println("Success")

  return
}
