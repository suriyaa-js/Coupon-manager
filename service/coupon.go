package coupon

import (
	"fmt"

	"suriyaa.com/coupon-manager/database"
	"suriyaa.com/coupon-manager/models"
	"suriyaa.com/coupon-manager/utils"
)

type couponService struct {
	db database.Database
}

type CouponService interface {
	GetCoupons() ([]models.Coupon, error)
	AddCoupon(coupon models.Coupon) (*models.Coupon, error)
	GetCouponByID(id int) (*models.Coupon, error)
	DeleteCouponByID(id int) (*models.Coupon, error)
	UpdateCouponByID(id int, coupon models.Coupon) (*models.Coupon, error)
	AddProduct(product models.Product) (*models.Product, error)
	DeleteProduct(productID int) (*models.Product, error)
}

func NewCouponService(db database.Database) CouponService {
	return &couponService{
		db: db,
	}
}

func (c *couponService) GetCoupons() ([]models.Coupon, error) {
	//fmt.Println("GetCoupon Service: Does nothing")
	coupons, err := c.db.GetCartCoupons()
	return coupons, err
}

func (c *couponService) AddCoupon(coupon models.Coupon) (*models.Coupon, error) {
	insertedCoupon, err := c.db.AddCoupon(coupon)
	if err != nil {
		fmt.Println("Error adding coupon: ", err)
		return nil, err
	}
	return insertedCoupon, nil
}

func (c *couponService) GetCouponByID(id int) (*models.Coupon, error) {
	coupon, err := c.db.GetCouponByID(id)
	if err != nil {
		fmt.Println("Error getting coupon by ID: ", err)
		return nil, err
	}
	return coupon, nil
}

// func (c *couponService) GetCouponByCode(code string) (*models.Coupon, error) {
// 	coupon, err := c.db.GetCouponByCode(code)
// 	if err != nil {
// 		fmt.Println("Error getting coupon by Code: ", err)
// 		return nil, err
// 	}
// 	return coupon, nil
// }

func (c *couponService) DeleteCouponByID(id int) (*models.Coupon, error) {
	coupon, err := c.db.DeleteCouponByID(id)
	if err != nil {
		fmt.Println("Error deleting coupon by ID", err)
		return nil, err
	}
	return coupon, nil
}

func (c *couponService) UpdateCouponByID(id int, coupon models.Coupon) (*models.Coupon, error) {
	_, err := utils.ValidateCouponType(coupon)
	if err != nil {
		fmt.Println("Error updating coupon by ID", err)
		return nil, err
	}
	updateCoupon, err := c.db.UpdateCouponByID(id, coupon)
	if err != nil {
		fmt.Println("Error updating coupon by ID", err)
		return nil, err
	}
	return updateCoupon, nil
}

func (c *couponService) AddProduct(product models.Product) (*models.Product, error) {
	coupon, err := c.db.AddProduct(product)
	if err != nil {
		fmt.Println("Error adding product to coupon", err)
		return nil, err
	}
	return coupon, nil
}

func (c *couponService) DeleteProduct(productID int) (*models.Product, error) {

	product, err := c.db.DeleteProduct(productID)
	if err != nil {
		fmt.Println("Error deleting product from coupon", err)
		return nil, err
	}
	return product, nil
}
