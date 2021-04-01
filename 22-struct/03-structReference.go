package main

import "fmt"

type Rectangle struct {
  width, height int
}

func main() {
  var rect1 Rectangle
  var rect2 *Rectangle = new(Rectangle)

  rect1.height = 28
  rect2.height = 62

  fmt.Println(rect1)
  fmt.Println(rect2)
}
