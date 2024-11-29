package repo

import (
	"fmt"

	"github.com/1Panel-dev/1Panel/core/constant"
	"github.com/1Panel-dev/1Panel/core/global"
	"gorm.io/gorm"
)

type ICommonRepo interface {
	WithByID(id uint) global.DBOption
	WithByGroupID(id uint) global.DBOption

	WithByName(name string) global.DBOption
	WithByType(ty string) global.DBOption
	WithByKey(key string) global.DBOption
	WithOrderBy(orderStr string) global.DBOption
	WithByStatus(status string) global.DBOption
	WithByGroupBelong(group string) global.DBOption

	WithByIDs(ids []uint) global.DBOption

	WithOrderRuleBy(orderBy, order string) global.DBOption
}

type CommonRepo struct{}

func NewICommonRepo() ICommonRepo {
	return &CommonRepo{}
}

func (c *CommonRepo) WithByID(id uint) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("id = ?", id)
	}
}
func (c *CommonRepo) WithByGroupID(id uint) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("group_id = ?", id)
	}
}

func (c *CommonRepo) WithByIDs(ids []uint) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("id in (?)", ids)
	}
}
func (c *CommonRepo) WithByName(name string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("`name` = ?", name)
	}
}

func (c *CommonRepo) WithByType(ty string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("`type` = ?", ty)
	}
}
func (c *CommonRepo) WithByKey(key string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("key = ?", key)
	}
}
func (c *CommonRepo) WithByStatus(status string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("status = ?", status)
	}
}
func (c *CommonRepo) WithByGroupBelong(group string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("group_belong = ?", group)
	}
}

func (c *CommonRepo) WithOrderBy(orderStr string) global.DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Order(orderStr)
	}
}

func (c *CommonRepo) WithOrderRuleBy(orderBy, order string) global.DBOption {
	switch order {
	case constant.OrderDesc:
		order = "desc"
	case constant.OrderAsc:
		order = "asc"
	default:
		orderBy = "created_at"
		order = "desc"
	}
	return func(g *gorm.DB) *gorm.DB {
		return g.Order(fmt.Sprintf("%s %s", orderBy, order))
	}
}
