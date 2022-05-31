package graphqlhandler

import (
	"kuncie/internal/modules/items/domain"
	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
)

// CommonFilter basic filter model
type CommonFilter struct {
	Limit   *int
	Page    *int
	Search  *string
	Sort    *string
	ShowAll *bool
	OrderBy *string
}

// toSharedFilter method
func (f *CommonFilter) toSharedFilter() (filter domain.FilterItems) {
	filter.Search = candihelper.PtrToString(f.Search)
	filter.OrderBy = candihelper.PtrToString(f.OrderBy)
	filter.Sort = candihelper.PtrToString(f.Sort)
	filter.ShowAll = candihelper.PtrToBool(f.ShowAll)

	if f.Limit == nil {
		filter.Limit = 10
	} else {
		filter.Limit = *f.Limit
	}
	if f.Page == nil {
		filter.Page = 1
	} else {
		filter.Page = *f.Page
	}

	return
}

// ItemsListResolver resolver
type ItemsListResolver struct {
	Meta candishared.Meta
	Data []shareddomain.Items
}
