package main

import "fmt"

type Rectangle struct {
  width, height int
}

func NewRectangle(width, height int) *Rectangle {
  return &Rectangle{width, height}
}

func main() {
  rect1 := NewRectangle(20, 10)
  rect2 := NewRectangle(30, 9)
  rect3 := NewRectangle(40, 8)

  fmt.Println(rect1)
  fmt.Println(rect2)
  fmt.Println(rect3)
}
