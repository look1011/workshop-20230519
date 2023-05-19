package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer to:", i)
		defer func() {
			fmt.Println("defer func to:", i)
		}()
	}

	for i := 0; i < 5; i++ {
		defer fmt.Println("to:", i)
		func(val int) {
			defer fmt.Println("func to:", i)
		}(i)
	}
}
