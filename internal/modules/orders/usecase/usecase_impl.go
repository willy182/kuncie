//

package usecase

import (
	"context"
	"errors"
	"fmt"

	"kuncie/internal/modules/orders/domain"
	productDomain "kuncie/internal/modules/products/domain"
	"kuncie/pkg/helper"
	shareddomain "kuncie/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *ordersUsecaseImpl) GetAllOrders(ctx context.Context, filter *domain.FilterOrders) (data []shareddomain.Orders, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersUsecase:GetAllOrders")
	defer trace.Finish()

	data, err = uc.repoSQL.OrdersRepo().FetchAll(ctx, filter)
	if err != nil {
		return data, meta, err
	}
	count := uc.repoSQL.OrdersRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	return
}

func (uc *ordersUsecaseImpl) GetDetailOrders(ctx context.Context, id int) (data shareddomain.Orders, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersUsecase:GetDetailOrders")
	defer trace.Finish()

	repoFilter := domain.FilterOrders{ID: id}
	data, err = uc.repoSQL.OrdersRepo().Find(ctx, &repoFilter)
	return
}

func (uc *ordersUsecaseImpl) CreateOrders(ctx context.Context, data *shareddomain.OrdersInputResolver) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrdersUsecase:CreateOrders")
	defer trace.Finish()

	orderID := helper.GenerateOrderNo()

	err = uc.repoSQL.WithTransaction(ctx, func(txCtx context.Context) error {
		var total float64

		for _, val := range data.Cart {
			promos, e := uc.sharedUsecase.Promotions(ctx, val.Sku, val.Qty)
			if e != nil {
				return e
			}

			for _, promo := range promos {
				// cek stock
				filterProduct := &productDomain.FilterProducts{Sku: promo.Sku}
				product, err := uc.repoSQL.ProductsRepo().Find(txCtx, filterProduct)
				if err != nil {
					return err
				}

				if product.Stock < promo.Stock {
					return errors.New(fmt.Sprintf("insufficient stock for SKU %s", promo.Sku))
				}

				payloadItems := &shareddomain.Items{
					OrderID: orderID,
					Sku:     promo.Sku,
					Qty:     promo.Stock,
					Amount:  promo.Price,
				}

				// sum total
				total += promo.Price

				if promo.Stock > 0 {
					// insert data items
					err = uc.repoSQL.ItemsRepo().Save(txCtx, payloadItems)
					if err != nil {
						return err
					}
				}

				// update data product
				product.Stock = product.Stock - promo.Stock
				err = uc.sharedUsecase.UpdateProducts(txCtx, promo.Sku, &product)
				if err != nil {
					return err
				}
			}

			// insert data orders
			payloadOrder := &shareddomain.Orders{
				OrderID: orderID,
				Total:   total,
			}
			err = uc.repoSQL.OrdersRepo().Save(txCtx, payloadOrder)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}
