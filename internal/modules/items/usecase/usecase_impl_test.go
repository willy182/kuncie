//

package usecase

import (
	"context"

	mockrepo "kuncie/pkg/mocks/modules/items/repository"
	mocksharedrepo "kuncie/pkg/mocks/shared/repository"
	shareddomain "kuncie/pkg/shared/domain"
	"testing"

	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewItemsUsecase(t *testing.T) {
	mockCache := &mockinterfaces.Cache{}
	mockRedisPool := &mockinterfaces.RedisPool{}
	mockRedisPool.On("Cache").Return(mockCache)

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetRedisPool").Return(mockRedisPool)

	uc, setFunc := NewItemsUsecase(mockDeps)
	setFunc(nil)
	assert.NotNil(t, uc)
}

func Test_itemsUsecaseImpl_CreateItems(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		itemsRepo := &mockrepo.ItemsRepository{}
		itemsRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ItemsRepo").Return(itemsRepo)

		uc := itemsUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.CreateItems(context.Background(), &shareddomain.Items{})
		assert.NoError(t, err)
	})
}
