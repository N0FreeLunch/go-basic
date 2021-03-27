package main

import "fmt"

func main () {
  a := []int{1,2,3,4,5}
  fmt.Println(len(a), cap(a))

  a = append(a, 6, 7)
  fmt.Println(len(a), cap(a))

  a = append(a, 8, 9, 10, 11)
  fmt.Println(len(a), cap(a))

  a = append(a, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
  fmt.Println(len(a), cap(a))
}
