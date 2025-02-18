package database

import (
	"fmt"
	"time"

	"suriyaa.com/coupon-manager/models"
)

func (d *database) AddProduct(product models.Product) (*models.Product, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	query := `INSERT INTO products (name, price, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, price, created_at, updated_at`
	var addedProduct models.Product
	err := d.db.QueryRow(query, product.Name, product.Price, product.CreatedAt, product.UpdatedAt).Scan(&addedProduct.ID, &addedProduct.Name, &addedProduct.Price, &addedProduct.CreatedAt, &addedProduct.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error inserting product: %w", err)
	}
	return &addedProduct, nil
}

func (d *database) DeleteProduct(productID int) (*models.Product, error) {
	query := `DELETE FROM products WHERE id = $1 RETURNING id, name, price, created_at, updated_at`
	var deletedProduct models.Product
	err := d.db.QueryRow(query, productID).Scan(&deletedProduct.ID, &deletedProduct.Name, &deletedProduct.Price, &deletedProduct.CreatedAt, &deletedProduct.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error deleting product: %w", err)
	}
	return &deletedProduct, nil
}
