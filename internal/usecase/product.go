package usecase

import (
	"context"

	"github.com/vanyovan/test-product.git/internal/entity"
	"github.com/vanyovan/test-product.git/internal/repo"
)

type ProductService struct {
	ProductRepo repo.ProductRepo
}

type ProductServiceProvider interface {
	CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error)
	UpdateProduct(ctx context.Context, id int64, product entity.Product) (err error)
	ViewProduct(ctx context.Context) (result []entity.Product, err error)
	DeleteProduct(ctx context.Context, id int64) (err error)
}

func NewProductService(ProductRepo repo.ProductRepo) ProductService {
	return ProductService{
		ProductRepo: ProductRepo,
	}
}

func (uc *ProductService) CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error) {
	result, err = uc.ProductRepo.CreateProduct(ctx, product)

	return result, err
}

func (uc *ProductService) UpdateProduct(ctx context.Context, id int64, product entity.Product) error {
	err := uc.ProductRepo.UpdateProductByProductID(ctx, id, product)

	return err
}

func (uc *ProductService) ViewProduct(ctx context.Context) (result []entity.Product, err error) {
	result, err = uc.ProductRepo.GetProducts(ctx)

	return result, err
}

func (uc *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	err := uc.ProductRepo.DeleteProductByProductID(ctx, id)

	return err
}
