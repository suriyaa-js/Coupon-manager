package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"suriyaa.com/coupon-manager/models"
	"suriyaa.com/coupon-manager/utils"
)

const (
	Conflict = "coupon already exists"
)

type Database interface {
	Migrate() error
	GetCartCoupons() ([]models.Coupon, error)
	AddCoupon(coupon models.Coupon) (*models.Coupon, error)
	GetCouponByID(id int) (*models.Coupon, error)
	DeleteCouponByID(id int) (*models.Coupon, error)
	UpdateCouponByID(id int, coupon models.Coupon) (*models.Coupon, error)
	AddProduct(product models.Product) (*models.Product, error)
	DeleteProduct(productID int) (*models.Product, error)
}

type database struct {
	db        *sql.DB
	isMigrate bool
}

// NewDatabase creates a new database instance
func NewDatabase(db *sql.DB, migrate bool) Database {
	return &database{
		db:        db,
		isMigrate: migrate,
	}
}

func (d *database) AddCoupon(coupon models.Coupon) (*models.Coupon, error) {
	// if coupon.CreatedAt == "" {
	coupon.CreatedAt = time.Now().Format(time.RFC3339)
	// }
	// if coupon.UpdatedAt == "" {
	coupon.UpdatedAt = time.Now().Format(time.RFC3339)
	// }

	id := utils.GenerateIDFromCode(coupon.Code)
	fmt.Println("Generated ID: ", id)
	coupon.ID = id

	if isExistAlready(id, d.db) {
		return nil, errors.New(Conflict)
	}

	query := `
        INSERT INTO coupons (id, code, type, discount_details, valid_from, valid_until, usage_limit, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := d.db.Exec(query, coupon.ID, coupon.Code, coupon.Type, coupon.DiscountDetails, coupon.ValidFrom, coupon.ValidUntil, coupon.UsageLimit, coupon.CreatedAt, coupon.UpdatedAt)
	if err != nil {
		fmt.Println("Error inserting coupon: ", err)
		return nil, fmt.Errorf("error inserting coupon: %w", err)
	}

	insertedCoupon, err := d.GetCouponByID(coupon.ID)
	if err != nil {
		return nil, err
	}

	return insertedCoupon, err
}

// GetCouponByID retrieves a coupon by its ID from the database
func (d *database) GetCouponByID(id int) (*models.Coupon, error) {
	query := `SELECT id, code, type, discount_details, valid_from, valid_until, usage_limit, created_at, updated_at FROM coupons WHERE id=$1`
	var coupon models.Coupon
	err := d.db.QueryRow(query, id).Scan(&coupon.ID, &coupon.Code, &coupon.Type, &coupon.DiscountDetails, &coupon.ValidFrom, &coupon.ValidUntil, &coupon.UsageLimit, &coupon.CreatedAt, &coupon.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("coupon with id %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving coupon: %w", err)
	}
	return &coupon, nil
}

func (d *database) GetCartCoupons() ([]models.Coupon, error) {
	query := `SELECT id, code, type, discount_details, valid_from, valid_until, usage_limit, created_at, updated_at FROM coupons`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving coupons: %w", err)
	}
	defer rows.Close()

	var coupons []models.Coupon
	for rows.Next() {
		var coupon models.Coupon
		err := rows.Scan(&coupon.ID, &coupon.Code, &coupon.Type, &coupon.DiscountDetails, &coupon.ValidFrom, &coupon.ValidUntil, &coupon.UsageLimit, &coupon.CreatedAt, &coupon.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning coupon: %w", err)
		}
		coupons = append(coupons, coupon)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return coupons, nil
}

func (d *database) DeleteCouponByID(id int) (*models.Coupon, error) {

	if !isExistAlready(id, d.db) {
		return nil, fmt.Errorf("coupon with id %d not found", id)
	}

	query := `DELETE FROM coupons WHERE id=$1`

	coupon, err := d.GetCouponByID(id)
	if err != nil {
		return nil, err
	}
	_, err = d.db.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("error deleting coupon: %w", err)
	}
	return coupon, nil
}

// isExistAlready Check if coupon already exists
func isExistAlready(id int, db *sql.DB) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM coupons WHERE id=$1)`
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking if coupon exists: ", err)
		return false
	}
	return exists
}

func (d *database) UpdateCouponByID(id int, coupon models.Coupon) (*models.Coupon, error) {
	existingCoupon, err := d.GetCouponByID(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving existing coupon: %w", err)
	}
	fmt.Println("Existing Coupon until Update: ", existingCoupon.ID)

	// Update only the fields that have non-empty values in the provided coupon parameter
	if coupon.Code != "" {
		existingCoupon.Code = coupon.Code
	}
	if coupon.Type != "" {
		existingCoupon.Type = coupon.Type
	}
	if len(coupon.DiscountDetails) > 0 {
		existingCoupon.DiscountDetails = coupon.DiscountDetails
	}
	if coupon.ValidFrom != "" {
		existingCoupon.ValidFrom = coupon.ValidFrom
	}
	if coupon.ValidUntil != "" {
		existingCoupon.ValidUntil = coupon.ValidUntil
	}
	if coupon.UsageLimit != 0 {
		existingCoupon.UsageLimit = coupon.UsageLimit
	}
	existingCoupon.UpdatedAt = time.Now().Format(time.RFC3339)

	// Update the coupon in the database
	query := `UPDATE coupons SET code=$1, type=$2, discount_details=$3, valid_from=$4, valid_until=$5, usage_limit=$6, updated_at=$7 WHERE id=$8`
	_, err = d.db.Exec(query, existingCoupon.Code, existingCoupon.Type, existingCoupon.DiscountDetails, existingCoupon.ValidFrom, existingCoupon.ValidUntil, existingCoupon.UsageLimit, existingCoupon.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("error updating coupon: %w", err)
	}

	return d.GetCouponByID(id)
}
