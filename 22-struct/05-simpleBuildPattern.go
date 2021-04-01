package main

import "fmt"

type Rectangle struct {
  width, height int
}

func main () {
  rect1 := &Rectangle{20, 10}
  rect2 := &Rectangle{30, 9}
  rect3 := &Rectangle{40, 8}

  fmt.Println(rect1)
  fmt.Println(rect2)
  fmt.Println(rect3)
}
