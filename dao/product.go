package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// GetProductById 通过 id 获取product
func (dao *ProductDao) GetProductById(id uint) (product model.Product, err error) {
	err = dao.Model(&model.Product{}).Where("id=?", id).
		First(&product).Error
	return
}

func (dao *ProductDao) ListProduct(id int) (products []model.Product, err error) {
	return
}
