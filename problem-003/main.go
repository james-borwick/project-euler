package main

import "fmt"

func main() {
	target := 13195

	for i := 1; i < target; i++ {
		if target%i == 0 {
			fmt.Println(i)
		}
	}
}
