package models

import (
	"encoding/json"
)

type CouponType string

const (
	CartWise    CouponType = "cart_wise"
	ProductWise CouponType = "product_wise"
	BxGy        CouponType = "bxgy"
)

// BaseCoupon represents the common fields for all coupon types
type BaseCoupon struct {
	ID         int        `json:"id"`
	Code       string     `json:"code"`
	Type       CouponType `json:"type"` // "cart-wise", "product-wise", "bxgy"
	ValidFrom  string     `json:"valid_from"`
	ValidUntil string     `json:"valid_until"`
	UsageLimit int        `json:"usage_limit"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

// Coupon represents the complete coupon structure
type Coupon struct {
	BaseCoupon
	DiscountDetails json.RawMessage `json:"discount_details"` // Raw JSON to be parsed based on type
	// Conditions      json.RawMessage `json:"conditions"`       // Raw JSON to be parsed based on type
}

// ParseDetails parses the DiscountDetails and Conditions based on the coupon type
// func (c *Coupon) ParseDetails() (interface{}, error) {
// 	switch c.Type {
// 	case "cart-wise":
// 		var details CartWiseDetails
// 		if err := json.Unmarshal(c.DiscountDetails, &details); err != nil {
// 			return nil, err
// 		}
// 		return details, nil
// 	case "product-wise":
// 		var details ProductWiseDetails
// 		if err := json.Unmarshal(c.DiscountDetails, &details); err != nil {
// 			return nil, err
// 		}
// 		return details, nil
// 	case "bxgy":
// 		var details BxGyDetails
// 		if err := json.Unmarshal(c.DiscountDetails, &details); err != nil {
// 			return nil, err
// 		}
// 		return details, nil
// 	default:
// 		return nil, fmt.Errorf("unknown coupon type: %s", c.Type)
// 	}
// }

// func _() {
// 	// Example JSON for a cart-wise coupon
// 	jsonData := `{
// 		"id": 1,
// 		"code": "CART10",
// 		"type": "cart-wise",
// 		"valid_from": "2023-10-01",
// 		"valid_until": "2023-10-31",
// 		"usage_limit": 100,
// 		"created_at": "2023-09-30",
// 		"updated_at": "2023-09-30",
// 		"discount_details": {
// 			"discount_type": "percentage",
// 			"value": 10
// 		},
// 		"conditions": {
// 			"min_cart_value": 100
// 		}
// 	}`

// 	var coupon Coupon
// 	if err := json.Unmarshal([]byte(jsonData), &coupon); err != nil {
// 		fmt.Println("Error unmarshaling JSON:", err)
// 		return
// 	}

// 	details, err := coupon.ParseDetails()
// 	bytes, _ := json.Marshal(details)
// 	coupon.DiscountDetails = json.RawMessage(bytes)
// 	if err != nil {
// 		fmt.Println("Error parsing details:", err)
// 		return
// 	}

// 	fmt.Printf("Coupon Details: %+v\n", details)
// }
