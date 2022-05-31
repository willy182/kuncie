// Code generated by candi v1.11.5.

package repository

import (
	"context"
	"time"

	"kuncie/internal/modules/products/domain"
	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"kuncie/pkg/shared"

	"gorm.io/gorm"
)

type productsRepoSQL struct {
	readDB, writeDB *gorm.DB
}

// NewProductsRepoSQL mongo repo constructor
func NewProductsRepoSQL(readDB, writeDB *gorm.DB) ProductsRepository {
	return &productsRepoSQL{
		readDB, writeDB,
	}
}

func (r *productsRepoSQL) Find(ctx context.Context, filter *domain.FilterProducts) (result shareddomain.Products, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProductsRepoSQL:Find")
	defer func() { trace.SetError(err); trace.Finish() }()

	db := shared.SetSpanToGorm(ctx, r.readDB)
	if filter.ID > 0 {
		db = db.Where("id = ?", filter.ID)
	}

	if filter.Sku != "" {
		db = db.Where("sku = ?", filter.Sku)
	}

	err = db.First(&result).Error
	return
}

func (r *productsRepoSQL) Save(ctx context.Context, data *shareddomain.Products) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProductsRepoSQL:Save")
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
