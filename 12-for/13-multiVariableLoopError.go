package main

import "fmt"

func main () {

  for i, j := 0, 0; i < 10; i++, j+=2 {
    fmt.Println(i, j)
  }

}
