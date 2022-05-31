// Code generated by candi v1.11.5.

package domain

import (
	"time"
)

// Products model
type Products struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:SERIAL"`
	Sku       string    `gorm:"column:sku;type:varchar(6)" json:"sku"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Price     float64   `gorm:"column:price" json:"price"`
	Stock     int       `gorm:"column:stock" json:"stock"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:TIMESTAMPTZ;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:TIMESTAMPTZ;default:current_timestamp"`
}

// Product model
type Product struct {
	Sku   string
	Name  string
	Price float64
	Qty   int
}

// TableName return table name of Products model
func (Products) TableName() string {
	return "products"
}