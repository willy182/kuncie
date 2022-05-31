// Code generated by candi v1.11.5.

package usecase

import (
	"sync"

	// @candi:usecaseImport
	itemsusecase "kuncie/internal/modules/items/usecase"
	ordersusecase "kuncie/internal/modules/orders/usecase"
	productsusecase "kuncie/internal/modules/products/usecase"
	"kuncie/pkg/shared/usecase/common"

	"github.com/golangid/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		// @candi:usecaseMethod
		Items() itemsusecase.ItemsUsecase
		Orders() ordersusecase.OrdersUsecase
		Products() productsusecase.ProductsUsecase
	}

	usecaseUow struct {
		// @candi:usecaseField
		itemsusecase.ItemsUsecase
		ordersusecase.OrdersUsecase
		productsusecase.ProductsUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = new(usecaseUow)
		var setSharedUsecaseFuncs []func(common.Usecase)
		var setSharedUsecaseFunc func(common.Usecase)

		// @candi:usecaseCommon
		usecaseInst.ItemsUsecase, setSharedUsecaseFunc = itemsusecase.NewItemsUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)
		usecaseInst.OrdersUsecase, setSharedUsecaseFunc = ordersusecase.NewOrdersUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)
		usecaseInst.ProductsUsecase, setSharedUsecaseFunc = productsusecase.NewProductsUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)

		sharedUsecase := common.SetCommonUsecase(usecaseInst)
		for _, setFunc := range setSharedUsecaseFuncs {
			setFunc(sharedUsecase)
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}

func (uc *usecaseUow) Items() itemsusecase.ItemsUsecase {
	return uc.ItemsUsecase
}

func (uc *usecaseUow) Orders() ordersusecase.OrdersUsecase {
	return uc.OrdersUsecase
}

func (uc *usecaseUow) Products() productsusecase.ProductsUsecase {
	return uc.ProductsUsecase
}
