package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"suriyaa.com/coupon-manager/models"
)

func ValidateAddProduct(product models.Product) (int, error) {
	// Check if required fields are present
	if product.Name == "" || product.Price == 0 {
		return http.StatusBadRequest, errors.New("invalid product details")
	}

	// Marshal the product to JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		return http.StatusInternalServerError, errors.New("error marshaling product")
	}

	// Unmarshal the JSON with DisallowUnknownFields to check for extra fields
	var validatedProduct models.Product
	decoder := json.NewDecoder(bytes.NewReader(productJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&validatedProduct); err != nil {
		return http.StatusBadRequest, errors.New("invalid product details: extra fields are not allowed")
	}

	return http.StatusOK, nil
}
