//

package usecase

import (
	"context"

	shareddomain "kuncie/pkg/shared/domain"
	"kuncie/pkg/shared/repository"
	"kuncie/pkg/shared/usecase/common"

	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
)

// ItemsUsecase abstraction
type ItemsUsecase interface {
	CreateItems(ctx context.Context, data *shareddomain.Items) (err error)
}

type itemsUsecaseImpl struct {
	sharedUsecase common.Usecase
	cache         interfaces.Cache
	repoSQL       repository.RepoSQL
}

// NewItemsUsecase usecase impl constructor
func NewItemsUsecase(deps dependency.Dependency) (ItemsUsecase, func(sharedUsecase common.Usecase)) {
	uc := &itemsUsecaseImpl{
		repoSQL: repository.GetSharedRepoSQL(),
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
