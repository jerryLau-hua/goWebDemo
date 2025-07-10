package service

import (
	"awesomeProject/internal/models"     // 确保这是你的模型包路径
	"awesomeProject/internal/repository" // 确保这是你的仓库包路径
	"context"
	"errors" // 导入errors包，用于创建自定义错误
)

// ProductService 定义了产品相关的业务逻辑接口
type ProductService interface {
	CreateProduct(ctx context.Context, name string, price float64, stock int) (*models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	UpdateProduct(ctx context.Context, id int64, name string, price float64, stock int) (*models.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}

// productService 是 ProductService 的具体实现
type productService struct {
	productRepo repository.ProductRepository // 它依赖于 ProductRepository 接口，而不是具体实现
}

// NewProductService 是 ProductService 的构造函数
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{productRepo: repo}
}

// CreateProduct 处理创建新产品的业务逻辑
func (s *productService) CreateProduct(ctx context.Context, name string, price float64, stock int) (*models.Product, error) {
	// 在这里可以添加业务逻辑，例如：
	// 1. 验证产品名称是否有效
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	// 2. 验证价格是否合法
	if price <= 0 {
		return nil, errors.New("product price must be positive")
	}
	if stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}
	// 创建一个新的产品模型实例
	product := &models.Product{
		Name:  name,
		Price: price, // 假设你的 Product 模型有 Price 字段
		Stock: stock,
	}

	// 调用仓库层来持久化数据
	err := s.productRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct 处理获取单个产品的业务逻辑
func (s *productService) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	// 直接调用仓库层。如果未来有缓存逻辑，可以加在这里。
	return s.productRepo.FindByID(ctx, id)
}

// GetAllProducts 处理获取所有产品的业务逻辑
func (s *productService) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	return s.productRepo.FindAll(ctx)
}

// UpdateProduct 处理更新产品的业务逻辑
func (s *productService) UpdateProduct(ctx context.Context, id int64, name string, price float64, stock int) (*models.Product, error) {
	// 1. 首先，获取要更新的产品
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err // 如果在查找过程中发生数据库错误
	}
	if product == nil {
		return nil, errors.New("product not found") // 如果产品不存在
	}

	// 2. 更新产品的字段
	product.Name = name
	product.Price = price
	product.Stock = stock

	// 3. 在这里可以添加更复杂的验证逻辑...

	// 4. 调用仓库层的更新方法
	err = s.productRepo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct 处理删除产品的业务逻辑
func (s *productService) DeleteProduct(ctx context.Context, id int64) error {
	// 在删除前，可以添加权限检查等业务逻辑
	// 例如：检查当前用户是否有权限删除该产品

	return s.productRepo.Delete(ctx, id)
}
