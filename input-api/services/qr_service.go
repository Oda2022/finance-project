package services

import (
	"input-api/models"

	"gonum.org/v1/gonum/mat"
)

func CalculateQR(matrix [][]float64) (*models.QrResponse, error) {
	rows := len(matrix)
	cols := len(matrix[0])

	data := make([]float64, 0, rows*cols)

	for _, row := range matrix {
		data = append(data, row...)
	}

	A := mat.NewDense(rows, cols, data)

	var qr mat.QR
	qr.Factorize(A)

	var q mat.Dense
	var r mat.Dense

	qr.QTo(&q)
	qr.RTo(&r)

	return &models.QrResponse{
		Q: denseToSlice(&q),
		R: denseToSlice(&r),
	}, nil
}

func denseToSlice(matrix *mat.Dense) [][]float64 {
	rows, cols := matrix.Dims()

	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)

		for j := 0; j < cols; j++ {
			result[i][j] = matrix.At(i, j)
		}
	}

	return result
}
