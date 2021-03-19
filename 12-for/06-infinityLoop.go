package main

import "fmt"
import "time"

func main () {

  i := 0
  for {
    fmt.Println(i)
    i = i + 1
    time.Sleep(500 * time.Millisecond)
  }

}
