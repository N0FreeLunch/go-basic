package main

import "fmt"
import "io/ioutil"

func main () {
  var b []byte
  var err error

  b, err = ioutil.ReadFile("./error.txt")

  if err == nil {
    fmt.Printf("%s", b)
  } else {
    fmt.Println("error")
  }
}
