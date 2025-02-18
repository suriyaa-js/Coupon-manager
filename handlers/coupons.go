package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"suriyaa.com/coupon-manager/models"
	coupon "suriyaa.com/coupon-manager/service"
	"suriyaa.com/coupon-manager/utils"
)

type handler struct {
	service coupon.CouponService
}

type API interface {
	http.Handler
	GetCoupons(writer http.ResponseWriter, req *http.Request)
	AddCoupon(writer http.ResponseWriter, req *http.Request)
	DeleteCoupon(writer http.ResponseWriter, req *http.Request)
	GetCouponByID(writer http.ResponseWriter, req *http.Request)
	UpdateCoupon(writer http.ResponseWriter, req *http.Request)
	AddProduct(writer http.ResponseWriter, req *http.Request)
	DeleteProduct(writer http.ResponseWriter, req *http.Request)
}

func NewHandler(service coupon.CouponService) API {
	return &handler{
		service: service,
	}
}

func (h *handler) AddCoupon(writer http.ResponseWriter, req *http.Request) {
	// Get cart coupon
	writer.Header().Set("Content-Type", "application/json")
	var coupon models.Coupon

	// Read the request body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// // Convert the body to a string and print it
	// bodyString := string(bodyBytes)
	// fmt.Println("Request Body:", bodyString)

	if err := json.Unmarshal(bodyBytes, &coupon); err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	statusCode, err := utils.ValidateBaseCoupon(coupon)
	if err != nil {
		returnError(writer, statusCode, err.Error())
		return
	}

	insertedCoupon, err := h.service.AddCoupon(coupon)
	if err != nil {
		returnError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(insertedCoupon)
}

func (h *handler) GetCoupons(writer http.ResponseWriter, req *http.Request) {
	// Get cart coupon
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Println("Get cart coupons handler called")
	coupons, err := h.service.GetCoupons()
	if err != nil {
		returnError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(writer).Encode(coupons)
}

func (h *handler) GetCouponByID(writer http.ResponseWriter, req *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	idStr := req.Context().Value("id")
	if idStr == "" {
		returnError(writer, http.StatusBadRequest, "Missing id query parameter")
		return
	}
	str := fmt.Sprintf("%v", idStr)
	fmt.Println("String value:", str)

	// Convert the id parameter to an integer
	id, err := strconv.Atoi(str)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid id query parameter")
		return
	}

	// Retrieve the coupon by ID
	coupon, err := h.service.GetCouponByID(id)
	if err != nil {
		returnError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupon)
}

func (h *handler) DeleteCoupon(writer http.ResponseWriter, req *http.Request) {
	// Get cart coupon
	writer.Header().Set("Content-Type", "application/json")
	idStr := req.Context().Value("id")
	if idStr == "" {
		returnError(writer, http.StatusBadRequest, "Missing id query parameter")
		return
	}
	str := fmt.Sprintf("%v", idStr)
	fmt.Println("String value:", str)

	// Convert the id parameter to an integer
	id, err := strconv.Atoi(str)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid id query parameter")
		return
	}
	coupon, err := h.service.DeleteCouponByID(id)
	if err != nil {
		returnError(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupon)
}

func (h *handler) UpdateCoupon(writer http.ResponseWriter, req *http.Request) {
	var coupon models.Coupon

	if err := json.NewDecoder(req.Body).Decode(&coupon); err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// statusCode, err := utils.ValidateBaseCoupon(coupon)
	// if err != nil {
	// 	returnError(writer, statusCode, err.Error())
	// 	return
	// }
	idStr := req.Context().Value("id")
	if idStr == "" {
		returnError(writer, http.StatusBadRequest, "Missing id query parameter")
		return
	}

	str := fmt.Sprintf("%v", idStr)
	fmt.Println("String value:", str)
	id, err := strconv.Atoi(str)
	if err != nil {
		returnError(writer, http.StatusBadRequest, "Invalid id query parameter")
		return
	}

	updatedCoupon, err := h.service.UpdateCouponByID(id, coupon)
	if err != nil {
		returnError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedCoupon)
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Println("Get cart coupons handler called")
	json.NewEncoder(writer).Encode("It does nothing")
}
