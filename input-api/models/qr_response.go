package models

type QrResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}
