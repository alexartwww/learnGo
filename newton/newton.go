package newton

import (
	"fmt"
	"math"
)

func Sqrt(params ...float64) float64 {
	var x, z, accuracy, delta float64

	if len(params) < 1 {
		panic("Sqrt must have at least 1 parameter")
	}

	x = params[0]
	z = params[0]
	delta = 1

	if len(params) == 2 {
		accuracy = params[1]
	} else {
		accuracy = 0.001
	}

	for ; delta > accuracy ; delta = ((math.Pow(z, 2) - x) / (2 * z)) {
		z = z - delta
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
