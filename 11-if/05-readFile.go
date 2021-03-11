package main

import "fmt"
import "io/ioutil"

func main () {
  var b []byte
  var err error

  if b, err = ioutil.ReadFile("./Hello.txt"); err == nil {
    fmt.Printf("%s", b)
  }

}
