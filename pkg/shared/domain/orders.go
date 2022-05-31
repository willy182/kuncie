//

package domain

import (
	"time"
)

// Orders model
type Orders struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:SERIAL"`
	OrderID   string    `gorm:"column:order_id;type:varchar(10)" json:"order_id"`
	Total     float64   `gorm:"column:total" json:"total"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:TIMESTAMPTZ;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:TIMESTAMPTZ;default:current_timestamp"`
}

// OrdersInputResolver model
type OrdersInputResolver struct {
	Cart []DataOrders
}

// DataOrders model
type DataOrders struct {
	Sku string
	Qty int
}

// TableName return table name of Orders model
func (Orders) TableName() string {
	return "orders"
}
