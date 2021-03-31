package main
import "fmt"

func factorial(n unit64) uint64 {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}

func main() {
  fmt.Println(factorial(5))
}
