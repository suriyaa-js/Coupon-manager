package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"suriyaa.com/coupon-manager/models"
	"suriyaa.com/coupon-manager/utils"
)

func (h *handler) AddProduct(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//var product models.Product

	// Read the request body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Decode the request body into a Product struct
	var product models.Product
	if err := json.Unmarshal(bodyBytes, &product); err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	code, err := utils.ValidateAddProduct(product)
	if err != nil {
		returnError(writer, code, err.Error())
		return
	}

	// Add the product using the service
	addedProduct, err := h.service.AddProduct(product)
	if err != nil {
		returnError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(addedProduct)
}

func (h *handler) DeleteProduct(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	idStr := req.Context().Value("id")
	if idStr == "" {
		returnError(writer, http.StatusBadRequest, "Missing id query parameter")
		return
	}
	str := fmt.Sprintf("%v", idStr)
	fmt.Println("String value:", str)

	// Convert the id parameter to an integer
	productID, err := strconv.Atoi(str)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid id query parameter")
		return
	}

	// Delete the product using the service
	product, err := h.service.DeleteProduct(productID)
	if err != nil {
		returnError(writer, http.StatusNotAcceptable, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
}
