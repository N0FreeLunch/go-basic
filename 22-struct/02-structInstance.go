package main

import "fmt"

type Rectangle struct {
  width, height int
}

func main () {

  var rect1 Rectangle = Rectangle{10, 20}
  rect2 := Rectangle{42, 62}
  rect3 := Rectangle{width : 30, height : 15}

  fmt.Println(rect1)
  fmt.Println(rect1.width)
  fmt.Println(rect1.height)
  fmt.Println(rect2)
  fmt.Println(rect2.width)
  fmt.Println(rect2.height)
  fmt.Println(rect3)
  fmt.Println(rect3.width)
  fmt.Println(rect3.height)
}
