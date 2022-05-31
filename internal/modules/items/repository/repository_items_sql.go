//

package repository

import (
	"context"

	"time"

	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"kuncie/pkg/shared"

	"gorm.io/gorm"
)

type itemsRepoSQL struct {
	readDB, writeDB *gorm.DB
}

// NewItemsRepoSQL mongo repo constructor
func NewItemsRepoSQL(readDB, writeDB *gorm.DB) ItemsRepository {
	return &itemsRepoSQL{
		readDB, writeDB,
	}
}

func (r *itemsRepoSQL) Save(ctx context.Context, data *shareddomain.Items) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ItemsRepoSQL:Save")
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
