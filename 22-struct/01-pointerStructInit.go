package main

import "fmt"

type Ractangle struct {
  width int
  height int
}

func main () {
  var rect1 *Ractangle
  rect1 = new(Ractangle)
  rect2 := new(Ractangle)

  fmt.Println(rect1)
  fmt.Println(rect2)
}
