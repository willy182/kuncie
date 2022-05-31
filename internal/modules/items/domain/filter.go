//

package domain

import "github.com/golangid/candi/candishared"

// FilterItems model
type FilterItems struct {
	candishared.Filter
	ID      int
	OrderID string `json:"order_id,omitempty"`
}
