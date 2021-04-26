package main

import (
  "fmt"
  "runtime"
)

func main() {
  cpuNumber:= runtime.NumCPU()
  fmt.Println("the number of computer cpu : ", cpuNumber)
  runtime.GOMAXPROCS(cpuNumber)
  toUseCPUNumer := runtime.GOMAXPROCS(0)
  fmt.Println("goroutine use cpu number : ", toUseCPUNumer)

  s := "Hello, world"

  for i := 0; i < 100; i++ {
    go func(n int) {
      fmt.Println("run function by goroutine : ", s, n)
    }(i)
  }

  fmt.Scanln()
}
