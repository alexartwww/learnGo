package tour_pic

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for i := 0; i < dx; i++ {
		row := make([]uint8, dy)
		result = append(result, row)
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i + j/2)
		}
	}
	return result
}
