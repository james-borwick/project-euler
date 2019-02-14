package main

import "fmt"

func main(){
	a := 1
	b := 2
	c := 0
	sum := 2

	for c <= 4000000 {
		c = a + b
		a = b
		b = c
		if c % 2 == 0 {
			sum += c
		}
		// fmt.Println()
		// fmt.Println("c is", c)
		// fmt.Println("sum is", sum)
		// fmt.Println(c)
	}
	fmt.Println("Sum of even numbers:", sum)
}