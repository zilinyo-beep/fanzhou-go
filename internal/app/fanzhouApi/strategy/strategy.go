package strategy

import (
	"context"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
)

type TaskStrategy interface {
	Execute(ctx context.Context, task *entity.Task) error
}

var strategyMap = map[string]TaskStrategy{
	"imagine":   &ImagineStrategy{},
	"blend":     &BlendStrategy{},
	"describe":  &DescribeStrategy{},
	"upscale":   &UpscaleStrategy{},
	"variation": &VariationStrategy{},
}

func GetStrategy(taskType string) TaskStrategy {
	if strategy, exists := strategyMap[taskType]; exists {
		return strategy
	}
	return nil
}
