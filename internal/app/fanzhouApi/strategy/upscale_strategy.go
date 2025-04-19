package strategy

import (
	"context"
	"fmt"
)

type UpscaleStrategy struct{}

func (s *UpscaleStrategy) Execute(ctx context.Context, task *entity.Task) error {
	fmt.Println("Upscale 执行中")
	return nil
}
