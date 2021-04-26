package main

import (
  "fmt"
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(1)
  useOneCoreCPU := runtime.GOMAXPROCS(0)
  fmt.Println("goroutine use cpu number : ", useOneCoreCPU)

  s := "Hello, world!"

  for i := 0; i < 100; i++ {
    go func () {
      fmt.Println(s, i)
    }()
  }

  fmt.Scanln()
}
