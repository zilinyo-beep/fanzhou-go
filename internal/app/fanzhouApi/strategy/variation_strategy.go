package strategy

import (
	"context"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"fmt"
)

type VariationStrategy struct{}

func (s *VariationStrategy) Execute(ctx context.Context, task *entity.Task) error {
	fmt.Println("Variation 执行中")
	return nil
}
