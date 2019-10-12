package repo

import (
	"context"
	"silkwormDemo/model"
	"silkwormDemo/internal/database"
	"github.com/jinzhu/gorm"
	"silkwormDemo/model/params"
)

type DemoRepo interface {
	MigrateTables() error

	GetDemoById(ctx context.Context, id uint64) *model.Demo
	ListDemo(ctx context.Context, req *params.ListDemoRequest) ([]*model.Demo, uint64, error)
	AddDemo(ctx context.Context, demo *model.Demo) error
	DeleteDemo(ctx context.Context, id uint64) error
}

type DemoRepoImpl struct{}

func (d *DemoRepoImpl) GetDB(ctx context.Context) *gorm.DB {
	return database.GetDB(ctx)
}

func (d *DemoRepoImpl) MigrateTables() error {
	db := d.GetDB(nil)
	return db.AutoMigrate(
		&model.Demo{},
	).Error
}

func (d *DemoRepoImpl) GetDemoById(ctx context.Context, id uint64) *model.Demo {
	var demo = new(model.Demo)
	db := d.GetDB(ctx)
	err := db.Where("id = ?", id).First(demo).Error
	if err != nil {
		return nil
	}
	return demo
}

func (d *DemoRepoImpl) ListDemo(ctx context.Context, req *params.ListDemoRequest) (
	demos []*model.Demo, total uint64, err error) {
	db := DefaultListDB(d.GetDB(ctx), &req.ListRequest)
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	err = db.Model(&model.Demo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if req.Limit != 0 {
		db = db.Limit(req.Limit)
	}
	if req.Offset != 0 {
		db = db.Offset(req.Offset)
	}
	err = db.Find(&demos).Error
	if err != nil {
		return nil, 0, err
	}
	return demos, total, nil
}

func (d *DemoRepoImpl) DeleteDemo(ctx context.Context, id uint64) error {
	return d.GetDB(ctx).Where(
		"id = ?", id,
	).Delete(model.Demo{}).Error
}

func (d *DemoRepoImpl) AddDemo(ctx context.Context, demo *model.Demo) error {
	return d.GetDB(ctx).Create(demo).Error
}
