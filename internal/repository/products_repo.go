package repository

import (
	"awesomeProject/internal/models"
	"context"
	"gorm.io/gorm"
)

// ProductRepository 定义了产品数据的所有操作，便于解耦
type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	FindAll(ctx context.Context) ([]*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int64) error
}

// 构造方法
type gormMySqlProductRepository struct {
	db *gorm.DB // 应用的是grom 框架 这里持有的是 *gorm.DB 而不是 *sql.DB
}

func (g gormMySqlProductRepository) Create(ctx context.Context, product *models.Product) error {
	result := g.db.WithContext(ctx).Create(product)
	return result.Error
}

func (g gormMySqlProductRepository) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	var product models.Product

	result := g.db.WithContext(ctx).First(&product, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &product, nil

}

func (g gormMySqlProductRepository) FindAll(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	result := g.db.WithContext(ctx).Find(&products)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return products, nil
}

func (g gormMySqlProductRepository) Update(ctx context.Context, product *models.Product) error {
	result := g.db.WithContext(ctx).Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g gormMySqlProductRepository) Delete(ctx context.Context, id int64) error {
	result := g.db.WithContext(ctx).Delete(&models.Product{}, id)
	return result.Error
}

// NewProductRepository 创建一个新的 ProductRepository 实例
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &gormMySqlProductRepository{db: db}
}
