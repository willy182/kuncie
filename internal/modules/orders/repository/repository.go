//

package repository

import (
	"context"

	"kuncie/internal/modules/orders/domain"
	shareddomain "kuncie/pkg/shared/domain"
)

// OrdersRepository abstract interface
type OrdersRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterOrders) ([]shareddomain.Orders, error)
	Count(ctx context.Context, filter *domain.FilterOrders) int
	Find(ctx context.Context, filter *domain.FilterOrders) (shareddomain.Orders, error)
	Save(ctx context.Context, data *shareddomain.Orders) error
}
