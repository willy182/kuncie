//

package graphqlhandler

import (
	"context"

	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

type mutationResolver struct {
	root *GraphQLHandler
}

// CreateOrders resolver
func (m *mutationResolver) CreateOrders(ctx context.Context, input struct {
	Data shareddomain.OrdersInputResolver
}) (ok string, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersDeliveryGraphQL:CreateOrders")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if err := m.root.uc.Orders().CreateOrders(ctx, &input.Data); err != nil {
		return ok, err
	}
	return "Success", nil
}
