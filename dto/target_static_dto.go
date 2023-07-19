package dto

import (
	"github/adekang/gin-demo/model"
	"strings"
)

type TargetStaticDto struct {
	ID           int    `json:"id"`
	ExpId        int    `json:"exp_id"`
	Connector    string `json:"connector"`
	Used         string `json:"used"`
	Value        string `json:"value"`
	TargetString string `json:"targetString"`
}

func ToTargetStaticDto(targetStatic []model.TargetStatic) []TargetStaticDto {

	var targetStaticDto []TargetStaticDto
	for _, v := range targetStatic {

		conn := strings.Split(v.Connector, ",")
		val := strings.Split(v.Value, ",")
		if len(conn) > 1 {
			for i, _ := range conn {
				targetStaticDto = append(targetStaticDto, TargetStaticDto{
					ID:           int(v.ID),
					ExpId:        v.ExpId,
					Connector:    conn[i],
					Used:         v.Used,
					Value:        val[i],
					TargetString: v.Expression.Name,
				})
			}
		} else {
			targetStaticDto = append(targetStaticDto, TargetStaticDto{
				ID:           int(v.ID),
				ExpId:        v.ExpId,
				Connector:    v.Connector,
				Used:         v.Used,
				Value:        v.Value,
				TargetString: v.Expression.Name,
			})
		}

	}
	return targetStaticDto
}
