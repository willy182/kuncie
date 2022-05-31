//

package graphqlhandler

import (
	"kuncie/pkg/shared/usecase"

	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/codebase/interfaces"
)

// GraphQLHandler model
type GraphQLHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewGraphQLHandler delivery
func NewGraphQLHandler(uc usecase.Usecase, deps dependency.Dependency) *GraphQLHandler {
	return &GraphQLHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// RegisterMiddleware register resolver based on schema in "api/graphql/*" path
func (h *GraphQLHandler) RegisterMiddleware(mwGroup *types.MiddlewareGroup) {
	mwGroup.Add("OrdersQueryResolver.getAllOrders", h.mw.GraphQLBearerAuth, h.mw.GraphQLPermissionACL("resource.public"))
	mwGroup.Add("OrdersQueryResolver.getDetailOrders", h.mw.GraphQLBearerAuth, h.mw.GraphQLPermissionACL("resource.public"))
	mwGroup.Add("OrdersMutationResolver.createOrders", h.mw.GraphQLBearerAuth, h.mw.GraphQLPermissionACL("resource.public"))
	// mwGroup.Add("OrdersMutationResolver.deleteOrders", h.mw.GraphQLBearerAuth, h.mw.GraphQLPermissionACL("resource.public"))
}

// Query method
func (h *GraphQLHandler) Query() interface{} {
	return &queryResolver{root: h}
}

// Mutation method
func (h *GraphQLHandler) Mutation() interface{} {
	return &mutationResolver{root: h}
}

// Subscription method
func (h *GraphQLHandler) Subscription() interface{} {
	return &subscriptionResolver{root: h}
}
