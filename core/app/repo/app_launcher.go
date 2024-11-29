package repo

import (
	"github.com/1Panel-dev/1Panel/core/app/model"
	"github.com/1Panel-dev/1Panel/core/global"
)

type LauncherRepo struct{}

type ILauncherRepo interface {
	Get(opts ...global.DBOption) (model.AppLauncher, error)
	List(opts ...global.DBOption) ([]model.AppLauncher, error)
	Create(launcher *model.AppLauncher) error
	Delete(opts ...global.DBOption) error
}

func NewILauncherRepo() ILauncherRepo {
	return &LauncherRepo{}
}

func (u *LauncherRepo) Get(opts ...global.DBOption) (model.AppLauncher, error) {
	var launcher model.AppLauncher
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(&launcher).Error
	return launcher, err
}
func (u *LauncherRepo) List(opts ...global.DBOption) ([]model.AppLauncher, error) {
	var ops []model.AppLauncher
	db := global.DB.Model(&model.AppLauncher{})
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(&ops).Error
	return ops, err
}

func (u *LauncherRepo) Create(launcher *model.AppLauncher) error {
	return global.DB.Create(launcher).Error
}

func (u *LauncherRepo) Delete(opts ...global.DBOption) error {
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	return db.Delete(&model.AppLauncher{}).Error
}
