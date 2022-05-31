//

package repository

import (
	"context"

	"time"

	"kuncie/internal/modules/orders/domain"
	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"kuncie/pkg/shared"

	"gorm.io/gorm"
)

type ordersRepoSQL struct {
	readDB, writeDB *gorm.DB
}

// NewOrdersRepoSQL mongo repo constructor
func NewOrdersRepoSQL(readDB, writeDB *gorm.DB) OrdersRepository {
	return &ordersRepoSQL{
		readDB, writeDB,
	}
}

func (r *ordersRepoSQL) FetchAll(ctx context.Context, filter *domain.FilterOrders) (data []shareddomain.Orders, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersRepoSQL:FetchAll")
	defer func() { trace.SetError(err); trace.Finish() }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}

	db := shared.SetSpanToGorm(ctx, r.readDB)

	err = db.Order(filter.OrderBy + " " + filter.Sort).
		Limit(filter.Limit).Offset(filter.CalculateOffset()).
		Find(&data).Error
	return
}

func (r *ordersRepoSQL) Count(ctx context.Context, filter *domain.FilterOrders) (count int) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersRepoSQL:Count")
	defer trace.Finish()

	db := shared.SetSpanToGorm(ctx, r.readDB)

	var total int64
	db.Model(&shareddomain.Orders{}).Count(&total)
	count = int(total)

	trace.Log("count", count)
	return
}

func (r *ordersRepoSQL) Find(ctx context.Context, filter *domain.FilterOrders) (result shareddomain.Orders, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersRepoSQL:Find")
	defer func() { trace.SetError(err); trace.Finish() }()

	db := shared.SetSpanToGorm(ctx, r.readDB)
	if filter.ID > 0 {
		db = db.Where("id = ?", filter.ID)
	}

	if filter.OrderID != "" {
		db = db.Where("order_id = ?", filter.OrderID)
	}

	err = db.First(&result).Error
	return
}

func (r *ordersRepoSQL) Save(ctx context.Context, data *shareddomain.Orders) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersRepoSQL:Save")
	defer func() { trace.SetError(err); trace.Finish() }()
	trace.Log("data", data)

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	data.UpdatedAt = time.Now()
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.ID > 0 {
		err = shared.SetSpanToGorm(ctx, db).Save(data).Error
	} else {
		err = shared.SetSpanToGorm(ctx, db).Create(data).Error
	}
	return
}
