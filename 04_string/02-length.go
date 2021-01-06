package main

import "fmt"
import "unicode/utf8"

func main () {
  var s1 string = "태스트"
  var s2 string = "test"

  fmt.Print("byte length : ")
  fmt.Println(len(s1))
  fmt.Print("byte length : ")
  fmt.Println(len(s2))
  fmt.Println(utf8.RuneCountInString(s1))
  fmt.Println(utf8.RuneCountInString(s2))

}
