package main

import "fmt"

func main() {
  a := [5]int{32,29,78,16,81}
  b := [5]int{
    32,
    29,
    78,
    16,
    81,
  }

  c := [...]int{32,29,78,16,81}
  d := [...]string{"Maria", "Andrew", "John"}

  // a := [3][3]int{
  //   {1,2,3},
  //   {4,5,6},
  //   {7,8,9},
  // }

  fmt.Println(a);
  fmt.Println(b);
  fmt.Println(c);
  fmt.Println(d);
}
