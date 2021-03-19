package main

import "fmt"
import "time"

func main () {

  i := 0
  for {
    if i > 4 {
      break;
    }
    fmt.Println(i)
    i = i + 1
    time.Sleep(500 * time.Millisecond)
  }

}
