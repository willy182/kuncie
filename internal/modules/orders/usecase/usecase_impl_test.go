//

package usecase

import (
	"context"
	"errors"

	"kuncie/internal/modules/orders/domain"
	mockrepo "kuncie/pkg/mocks/modules/orders/repository"
	mocksharedrepo "kuncie/pkg/mocks/shared/repository"
	shareddomain "kuncie/pkg/shared/domain"
	"testing"

	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewOrdersUsecase(t *testing.T) {
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

	uc, setFunc := NewOrdersUsecase(mockDeps)
	setFunc(nil)
	assert.NotNil(t, uc)
}

func Test_ordersUsecaseImpl_GetAllOrders(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		ordersRepo := &mockrepo.OrdersRepository{}
		ordersRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Orders{}, nil)
		ordersRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrdersRepo").Return(ordersRepo)

		uc := ordersUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllOrders(context.Background(), &domain.FilterOrders{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		ordersRepo := &mockrepo.OrdersRepository{}
		ordersRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Orders{}, errors.New("Error"))
		ordersRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrdersRepo").Return(ordersRepo)

		uc := ordersUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllOrders(context.Background(), &domain.FilterOrders{})
		assert.Error(t, err)
	})
}

func Test_ordersUsecaseImpl_GetDetailOrders(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		responseData := shareddomain.Orders{}

		ordersRepo := &mockrepo.OrdersRepository{}
		ordersRepo.On("Find", mock.Anything, mock.Anything).Return(responseData, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrdersRepo").Return(ordersRepo)

		uc := ordersUsecaseImpl{
			repoSQL: repoSQL,
		}

		result, err := uc.GetDetailOrders(context.Background(), 2)
		assert.NoError(t, err)
		assert.Equal(t, responseData, result)
	})
}

// func Test_ordersUsecaseImpl_CreateOrders(t *testing.T) {
// 	t.Run("Testcase #1: Positive", func(t *testing.T) {

// 		ordersRepo := &mockrepo.OrdersRepository{}
// 		ordersRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

// 		repoSQL := &mocksharedrepo.RepoSQL{}
// 		repoSQL.On("OrdersRepo").Return(ordersRepo)
// 		repoSQL.On("WithTransaction", mock.Anything, mock.Anything).Return(nil)

// 		uc := ordersUsecaseImpl{
// 			repoSQL: repoSQL,
// 		}

// 		err := uc.CreateOrders(context.Background(), &shareddomain.OrdersInputResolver{})
// 		assert.NoError(t, err)
// 	})
// }
