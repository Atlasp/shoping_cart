package model

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	OrderId       int64         `json:"order_id" gorm:"primaryKey"`
	CustomerId    int64         `json:"customer_id"`
	OrderedItems  pq.Int64Array `json:"ordered_items"  gorm:"type:integer[]"`
	OrderTotal    int64         `json:"order_total"`
	OrderDiscount int64         `json:"order_discount"`
	OrderFinal    int64         `json:"order_final"`
	OrderDate     time.Time     `json:"order_date"`
}
