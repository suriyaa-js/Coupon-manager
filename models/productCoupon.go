package models

// import (
// 	"encoding/json"
// 	"fmt"
// )

// ProductWiseDetails represents the discount_details and conditions for product-wise coupons
type ProductWiseDetails struct {
	DiscountType string  `json:"discount_type"` // "percentage" or "fixed_amount"
	Value        float64 `json:"value"`         // e.g., 20 for 20% or 15 for Rs. 15 off
	ProductIDs   []int   `json:"product_ids"`   // List of product IDs eligible for the discount
}
