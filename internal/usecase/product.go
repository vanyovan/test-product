package usecase

import (
	"context"
	"errors"

	"github.com/vanyovan/test-product.git/internal/entity"
	"github.com/vanyovan/test-product.git/internal/helper"
	"github.com/vanyovan/test-product.git/internal/repo"
)

type ProductService struct {
	ProductRepo repo.ProductRepo
}

type ProductServiceProvider interface {
	CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error)
	UpdateProduct(ctx context.Context) (result entity.Wallet, err error)
	ViewProduct(ctx context.Context) (result entity.Wallet, err error)
	DeleteProduct(ctx context.Context, id int64) (err error)
}

func NewProductService(ProductRepo repo.ProductRepo) ProductService {
	return ProductService{
		ProductRepo: ProductRepo,
	}
}

func (uc *ProductService) CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error) {
	result, err = uc.ProductRepo.CreateProduct(ctx, product)
	if err != nil {
		return result, err
	}
	return result, err
}

func (uc *ProductService) UpdateProduct(ctx context.Context) (result entity.Wallet, err error) {
	// get wallet.
	result, err = uc.ProductRepo.GetWalletByUserId(ctx, "productid")
	if helper.IsStructEmpty(result) || result.Status == helper.ConstantDisabled {
		return result, errors.New("wallet not found or wallet already disabled")
	}

	if err != nil {
		return result, err
	}

	//disable wallet
	updatedAt, err := uc.ProductRepo.UpdateStatusByUserId(ctx, helper.ConstantDisabled, "productid")
	if err != nil {
		return result, err
	}

	result.Status = helper.ConstantDisabled
	result.DisabledAt = &updatedAt

	return result, err
}

func (uc *ProductService) ViewProduct(ctx context.Context) (result []entity.Product, err error) {
	result, err = uc.ProductRepo.GetProducts(ctx)

	return result, err
}

func (uc *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	// get wallet.
	err := uc.ProductRepo.DeleteProductByProductID(ctx, id)

	return err
}
