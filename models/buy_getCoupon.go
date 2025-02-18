package models

// BxGyDetails represents the discount_details and conditions for BxGy coupons
type BxGyDetails struct {
	DiscountType    string         `json:"discount_type"` // "free", "percentage", or "fixed_amount"
	Value           float64        `json:"value"`         // e.g., 100 for 100% off (free)
	BuyProducts     []ProductGroup `json:"buy_products"`  // Products to buy
	GetProducts     []ProductGroup `json:"get_products"`  // Products to get free/discounted
	RepetitionLimit int            `json:"repetition_limit"`
}

// ProductGroup represents a group of products and their quantities
type ProductGroup struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
