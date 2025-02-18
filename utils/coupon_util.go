package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"suriyaa.com/coupon-manager/models"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ValidateBaseCoupon(coupon models.Coupon) (int, error) {
	// Validate coupon
	if coupon.Type != models.BxGy && coupon.Type != models.CartWise && coupon.Type != models.ProductWise {
		return http.StatusBadRequest, errors.New("invalid coupon type")
	}
	if coupon.DiscountDetails == nil {
		return http.StatusBadRequest, errors.New("invalid discount details")
	}
	if coupon.Code == "" {
		return http.StatusBadRequest, errors.New("invalid coupon code")
	}
	if coupon.DiscountDetails == nil {
		return http.StatusBadRequest, errors.New("invalid discount details")
	}
	return ValidateCouponType(coupon)
}

func ValidateCouponType(coupon models.Coupon) (int, error) {
	switch coupon.Type {
	case models.BxGy:
		return validateBxGyCoupon(coupon)
	case models.CartWise:
		return validateCartWiseCoupon(coupon)
	case models.ProductWise:
		return validateProductWiseCoupon(coupon)
	default:
		return http.StatusOK, nil
	}
}

func validateCartWiseCoupon(coupon models.Coupon) (int, error) {
	var details models.CartWiseDetails
	decoder := json.NewDecoder(bytes.NewReader(coupon.DiscountDetails))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&details); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid discount details: %w", err)
	}
	if details.DiscountType == "" || details.Value == 0 || details.MinCartValue == 0 {
		return http.StatusBadRequest, fmt.Errorf("all fields in discount details must have values for chosen coupon type")
	}
	return http.StatusOK, nil
}

func validateProductWiseCoupon(coupon models.Coupon) (int, error) {
	var details models.ProductWiseDetails
	decoder := json.NewDecoder(bytes.NewReader(coupon.DiscountDetails))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&details); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid discount details: %w", err)
	}
	if details.DiscountType == "" || details.Value == 0 || len(details.ProductIDs) == 0 {
		return http.StatusBadRequest, fmt.Errorf("all fields in discount details must have values for chosen coupon type")
	}
	return http.StatusOK, nil
}

func validateBxGyCoupon(coupon models.Coupon) (int, error) {
	var details models.BxGyDetails
	decoder := json.NewDecoder(bytes.NewReader(coupon.DiscountDetails))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&details); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid discount details: %w", err)
	}
	if details.DiscountType == "" || details.Value == 0 || len(details.BuyProducts) == 0 || len(details.GetProducts) == 0 {
		return http.StatusBadRequest, fmt.Errorf("all fields in discount details must have values for chosen coupon type except RepetitionLimit")
	}
	return http.StatusOK, nil
}

func GenerateIDFromCode(code string) int {
	// Convert the code to an integer using Base62 encoding
	fmt.Println("Code: ", code)
	return base62ToInt(code)
}

func GenerateCodeFromID(id int) string {
	// Convert the ID to a string using Base62 encoding
	return intToBase62(id)
}

func base62ToInt(s string) int {
	result := 0
	for _, c := range s {
		result = result*62 + indexOf(c)
	}
	return result
}

func intToBase62(n int) string {
	if n == 0 {
		return string(base62Chars[0])
	}
	result := ""
	for n > 0 {
		result = string(base62Chars[n%62]) + result
		n /= 62
	}
	return result
}

func indexOf(c rune) int {
	for i, char := range base62Chars {
		if char == c {
			return i
		}
	}
	return -1
}
