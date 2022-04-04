package domain

import (
	"product/utils/errs"
)

type ProductService interface {
	Register(Product) (*Product, *errs.AppError)
	FindById(string) (*Product, *errs.AppError)
	Update(string, Product) (*Product, *errs.AppError)
	DeleteById(string) (*Product, *errs.AppError)
}

type DefaultProductService struct {
	ProductDB ProductRepositoryDB
}

func NewProductService(productDB ProductRepositoryDB) ProductService {
	return &DefaultProductService{
		ProductDB: productDB,
	}
}

func (psvc DefaultProductService) Register(product Product) (*Product, *errs.AppError) {

	u, err := psvc.ProductDB.Save(product)
	return u, err
}

func (psvc *DefaultProductService) FindById(id string) (*Product, *errs.AppError) {

	product, err := psvc.ProductDB.FetchProductById(id)
	return product, err
}

func (psvc DefaultProductService) Update(id string, product Product) (*Product, *errs.AppError) {

	u, err := psvc.ProductDB.UpdateProduct(id, product)
	return u, err
}

func (psvc DefaultProductService) DeleteById(id string) (*Product, *errs.AppError) {

	u, err := psvc.FindById(id)
	if err != nil {
		return nil, err
	}
	err = psvc.ProductDB.DeleteProductById(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
