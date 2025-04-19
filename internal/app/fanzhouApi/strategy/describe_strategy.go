package strategy

import (
	"context"
	"fmt"
)

type DescribeStrategy struct{}

func (s *DescribeStrategy) Execute(ctx context.Context, task *entity.Task) error {
	fmt.Println("Describe 执行中")
	return nil
}
