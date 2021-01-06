package main

import "fmt"

func main () {
  var s1 string = "Hello, world\n"
  var s2 string = "안녕하세요\n"
  var s3 string = "\ud55c\uae00"
  var s4 string = "\U0000d55c\U0000ae00\n"
  var s5 string = "\xed\x95\x9c\xea\xb8\x80\n"
  var s6 string = "Hello, world\n"
  var s7 string = `
  Hello
  world
  `

  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(s3)
  fmt.Println(s4)
  fmt.Println(s5)
  fmt.Println(s6)
  fmt.Println(s7)

}
