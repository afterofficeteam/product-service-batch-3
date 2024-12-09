package mock_products

import (
	model "codebase-service/models"
	"codebase-service/repository/products"

	"github.com/stretchr/testify/mock"
)

var _ products.ProductRepository = &MockProductRepo{}

type MockProductRepo struct {
	mock.Mock
}

func NewMockProductRepo() *MockProductRepo {
	return &MockProductRepo{}
}

func (m *MockProductRepo) IsShopOwner(userId, shopId string) error {
	args := m.Called(userId, shopId)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockProductRepo) CreateProduct(req *model.CreateProductReq) (*model.GetProductResp, error) {
	args := m.Called(req)
	var (
		resp *model.GetProductResp
		err  error
	)

	if n, ok := args.Get(0).(*model.GetProductResp); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}
func (m *MockProductRepo) GetProduct(req *model.GetProductReq) (*model.GetProductResp, error) {
	args := m.Called(req)
	var (
		resp *model.GetProductResp
		err  error
	)

	if n, ok := args.Get(0).(*model.GetProductResp); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}

func (m *MockProductRepo) GetProducts(req *model.GetProductsReq) (*model.GetProductsResp, error) {
	args := m.Called(req)
	var (
		resp *model.GetProductsResp
		err  error
	)

	if n, ok := args.Get(0).(*model.GetProductsResp); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}

func (m *MockProductRepo) DeleteProduct(req *model.DeleteProductReq) error {
	args := m.Called(req)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
