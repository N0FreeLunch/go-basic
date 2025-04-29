package main

import "fmt"

func main() {
	var s string = "test"

	fmt.Print("文字のデフォルトフォマットは定数に表示: ")
	fmt.Println(s[1])
	fmt.Printf("%%cを利用して文字を表示: %c\n", s[1])
}
