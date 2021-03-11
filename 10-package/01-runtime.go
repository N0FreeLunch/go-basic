package main

import "fmt"
import "runtime"

func main() {
  fmt.Println("CPU count : ", runtime.NumCPU())
}
