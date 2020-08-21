package main

import (
	"fmt"
	"learnProject/tour_fibonacci"
)

func main() {
	f := tour_fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
