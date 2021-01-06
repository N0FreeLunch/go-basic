package main

import "fmt"

func main () {
  var s1 string = "Hello, world!\n"
  fmt.Printf("%c\n",s1[0])
  s1 = "abcdefg"
  fmt.Println(s1[0])
  fmt.Printf("%c\n",s1[0])
  // s1[0] = 'z' //error
}
