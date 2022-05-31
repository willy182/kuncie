//

package repository

import (
	"context"

	shareddomain "kuncie/pkg/shared/domain"
)

// ItemsRepository abstract interface
type ItemsRepository interface {
	Save(ctx context.Context, data *shareddomain.Items) error
}
