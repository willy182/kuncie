//

package usecase

import (
	"context"

	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *itemsUsecaseImpl) CreateItems(ctx context.Context, data *shareddomain.Items) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ItemsUsecase:CreateItems")
	defer trace.Finish()

	return uc.repoSQL.ItemsRepo().Save(ctx, data)
}
