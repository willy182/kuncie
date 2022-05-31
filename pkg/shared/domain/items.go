//

package domain

import (
	"time"
)

// Items model
type Items struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:SERIAL"`
	OrderID   string    `gorm:"column:order_id;type:varchar(10)" json:"order_id"`
	Sku       string    `gorm:"column:sku;type:varchar(6)" json:"sku"`
	Qty       int       `gorm:"column:qty;" json:"qty"`
	Amount    float64   `gorm:"column:amount" json:"amount"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:TIMESTAMPTZ;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:TIMESTAMPTZ;default:current_timestamp"`
}

// TableName return table name of Items model
func (Items) TableName() string {
	return "items"
}
