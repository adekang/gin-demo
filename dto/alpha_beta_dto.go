package dto

import "github/adekang/gin-demo/model"

type AlphaBetaDto struct {
	ID    int     `json:"id"`
	Alpha float32 `json:"alpha"`
	Beta  float32 `json:"beta"`
	Apply string  `json:"apply"`
}

func ToAlphaBetaDto(alphaBate []model.AlphaBeta) []AlphaBetaDto {

	var alphaBetaDto []AlphaBetaDto
	for _, v := range alphaBate {
		alphaBetaDto = append(alphaBetaDto, AlphaBetaDto{
			ID:    v.ID,
			Alpha: v.Alpha,
			Beta:  v.Beta,
			Apply: v.Apply,
		})
	}
	return alphaBetaDto

}
