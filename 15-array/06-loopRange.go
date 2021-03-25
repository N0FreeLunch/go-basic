package main

import "fmt"

func main () {
   a:= [5]int{32, 29, 78, 16, 81}

   for index := range a {
     fmt.Println(index)
   }

   b := [5]int{32, 29, 78, 16, 81}

   for _,element := range b {
     fmt.Println(element)
   }
}
