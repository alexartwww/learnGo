package main

import (
	"fmt"
	"learnProject/newton"
	"math"
)

func main() {
	fmt.Println(newton.Sqrt(10, 0.00001))
	fmt.Println(math.Sqrt(10))
}
