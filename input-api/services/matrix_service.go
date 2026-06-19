package services

func IsRectangular(matrix [][]float64) bool {
	if len(matrix) == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	for _, row := range matrix {
		if len(row) != cols {
			return false
		}
	}

	return true
}
