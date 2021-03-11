package main

import "fmt"

func main () {
  binding()
  line()
  defineVariable()
  line()
  subtraction()
  line()
  line()
  line()
  line()
  line()
  line()
  line()
  line()
}

func line () {
  fmt.Println("------")
}

func binding () {
  a := 1
  b := 2
  var c int = b
  const d string = "Hello, world"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

func defineVariable () {
  a := 1
  b := 3.5
  c := "Hello, world"
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

func subtraction() {
  a := 3 - 2
  b := 4 - 5
  c := a - b
  d := "Hello, " + "World!"
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
