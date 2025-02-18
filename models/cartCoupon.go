package models

// CartWiseDetails represents the discount_details and conditions for cart-wise coupons
type CartWiseDetails struct {
	DiscountType string  `json:"discount_type"` // "percentage" or "fixed_amount"
	Value        float64 `json:"value"`         // e.g., 10 for 10% or 20 for Rs. 20 off
	MinCartValue float64 `json:"min_cart_value"`
}
