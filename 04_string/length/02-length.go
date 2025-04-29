package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "태스트"
	var s2 string = "test"

	fmt.Printf("byte length of \"%s\" : ", s1)
	fmt.Println(len(s1))
	fmt.Printf("byte length of \"%s\" : ", s2)
	fmt.Println(len(s2))
	fmt.Printf("letter length of \"%s\" : ", s1)
	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Printf("letter length of \"%s\" : ", s2)
	fmt.Println(utf8.RuneCountInString(s2))

}
