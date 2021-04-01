package main

import "fmt"

type Rectangle struct {
  width, height int
}

func rectangleArea(rect *Rectangle) int {
  return rect.width * rect.height
}

func main () {
  rect := Rectangle{30, 30}
  area := rectangleArea(&rect)
  fmt.Println(area)
}
