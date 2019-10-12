package demo

import (
	"context"
	"silkwormDemo/model"
	"silkwormDemo/service/demo/cache"
)

type DemoService interface {
	GetDemoById(ctx context.Context, demoId uint64) (*model.Demo, error)
}

type DemoServiceImpl struct {
	Cache cache.DemoCache
}

