//

package usecase

import (
	"context"

	"kuncie/internal/modules/orders/domain"
	shareddomain "kuncie/pkg/shared/domain"
	"kuncie/pkg/shared/repository"
	"kuncie/pkg/shared/usecase/common"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
)

// OrdersUsecase abstraction
type OrdersUsecase interface {
	GetAllOrders(ctx context.Context, filter *domain.FilterOrders) (data []shareddomain.Orders, meta candishared.Meta, err error)
	GetDetailOrders(ctx context.Context, id int) (data shareddomain.Orders, err error)
	CreateOrders(ctx context.Context, data *shareddomain.OrdersInputResolver) (err error)
}

type ordersUsecaseImpl struct {
	sharedUsecase common.Usecase
	cache         interfaces.Cache
	repoSQL       repository.RepoSQL
	// repoMongo     repository.RepoMongo
	// kafkaPub      interfaces.Publisher
	// rabbitmqPub   interfaces.Publisher
}

// NewOrdersUsecase usecase impl constructor
func NewOrdersUsecase(deps dependency.Dependency) (OrdersUsecase, func(sharedUsecase common.Usecase)) {
	uc := &ordersUsecaseImpl{
		// cache: deps.GetRedisPool().Cache(),
		repoSQL: repository.GetSharedRepoSQL(),
		// repoMongo: repository.GetSharedRepoMongo(),
		// kafkaPub: deps.GetBroker(types.Kafka).GetPublisher(),
		// rabbitmqPub: deps.GetBroker(types.RabbitMQ).GetPublisher(),
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
