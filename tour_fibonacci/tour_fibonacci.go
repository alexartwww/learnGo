package tour_fibonacci

func Fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		result := i + j
		i = j
		j = result
		return result
	}
}
