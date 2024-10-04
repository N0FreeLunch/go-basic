package main

import "fmt"

func main() {
	var x, y int = 30, 50
	var age, name = 10, "Maria"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(age)
	fmt.Println(name)

	a, b, c, d := 1, 3.4, "Hello, world!", false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("-------")

	var (
		x_, y_      int = 30, 50
		age_, name_     = 10, "Maria"
	)

	fmt.Println(x_)
	fmt.Println(y_)
	fmt.Println(age_)
	fmt.Println(name_)

}
