//

package domain

import "github.com/golangid/candi/candishared"

// FilterOrders model
type FilterOrders struct {
	candishared.Filter
	ID      int
	OrderID string `json:"order_id,omitempty"`
}
