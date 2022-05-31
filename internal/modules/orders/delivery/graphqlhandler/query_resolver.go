//

package graphqlhandler

import (
	"context"

	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

type queryResolver struct {
	root *GraphQLHandler
}

// GetAllOrders resolver
func (q *queryResolver) GetAllOrders(ctx context.Context, input struct{ Filter *CommonFilter }) (results OrdersListResolver, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersDeliveryGraphQL:GetAllOrders")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if input.Filter == nil {
		input.Filter = new(CommonFilter)
	}
	filter := input.Filter.toSharedFilter()
	data, meta, err := q.root.uc.Orders().GetAllOrders(ctx, &filter)
	if err != nil {
		return results, err
	}

	return OrdersListResolver{
		Meta: meta, Data: data,
	}, nil
}

// GetDetailOrders resolver
func (q *queryResolver) GetDetailOrders(ctx context.Context, input struct{ ID int }) (data shareddomain.Orders, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersDeliveryGraphQL:GetDetailOrders")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	return q.root.uc.Orders().GetDetailOrders(ctx, input.ID)
}
