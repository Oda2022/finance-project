package services

func IsRectangular(matrix [][]float64) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	cols := len(matrix[0])

	for _, row := range matrix {
		if len(row) != cols {
			return false
		}
	}

	rows := len(matrix)
	return rows != cols
}