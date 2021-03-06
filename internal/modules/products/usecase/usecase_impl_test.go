// Code generated by candi v1.11.5.

package usecase

import (
	"context"

	mockrepo "kuncie/pkg/mocks/modules/products/repository"
	mocksharedrepo "kuncie/pkg/mocks/shared/repository"
	shareddomain "kuncie/pkg/shared/domain"
	"testing"

	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewProductsUsecase(t *testing.T) {
	/*
		mockPublisher := &mockinterfaces.Publisher{}
		mockBroker := &mockinterfaces.Broker{}
		mockBroker.On("GetPublisher").Return(mockPublisher)
	*/
	mockCache := &mockinterfaces.Cache{}
	mockRedisPool := &mockinterfaces.RedisPool{}
	mockRedisPool.On("Cache").Return(mockCache)

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetRedisPool").Return(mockRedisPool)
	// mockDeps.On("GetBroker", mock.Anything).Return(mockBroker)

	uc, setFunc := NewProductsUsecase(mockDeps)
	setFunc(nil)
	assert.NotNil(t, uc)
}

func Test_productsUsecaseImpl_GetDetailProducts(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		responseData := shareddomain.Products{}

		productsRepo := &mockrepo.ProductsRepository{}
		productsRepo.On("Find", mock.Anything, mock.Anything).Return(responseData, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProductsRepo").Return(productsRepo)

		uc := productsUsecaseImpl{
			repoSQL: repoSQL,
		}

		result, err := uc.GetDetailProducts(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, responseData, result)
	})
}

// func Test_ordersUsecaseImpl_UpdateProducts(t *testing.T) {
// 	t.Run("Testcase #1: Positive", func(t *testing.T) {

// 		ordersRepo := &mockrepo.ProductsRepository{}
// 		ordersRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Products{}, nil)
// 		ordersRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

// 		repoSQL := &mocksharedrepo.RepoSQL{}
// 		repoSQL.On("OrdersRepo").Return(ordersRepo)

// 		uc := productsUsecaseImpl{
// 			repoSQL: repoSQL,
// 		}

// 		err := uc.UpdateProducts(context.Background(), "sku1", &shareddomain.Products{})
// 		assert.NoError(t, err)
// 	})

// 	t.Run("Testcase #2: Negative", func(t *testing.T) {

// 		ordersRepo := &mockrepo.ProductsRepository{}
// 		ordersRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Products{}, errors.New("Error"))
// 		ordersRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

// 		repoSQL := &mocksharedrepo.RepoSQL{}
// 		repoSQL.On("OrdersRepo").Return(ordersRepo)

// 		uc := productsUsecaseImpl{
// 			repoSQL: repoSQL,
// 		}

// 		err := uc.UpdateProducts(context.Background(), "sku2", &shareddomain.Products{})
// 		assert.Error(t, err)
// 	})
// }
