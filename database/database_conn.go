package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"suriyaa.com/coupon-manager/server/props"
)

// ConnectDB connects to the database with the provided configuration
func ConnectDB(config props.DatabaseConfig) (*sql.DB, error) {
	// Connect to database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error: -connecting to database: ", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {

		return nil, fmt.Errorf("error pinging database: %w", err)
	}
	return db, nil
}

// Migrate runs the migrations if isMigrate flag is set to true
func (d *database) Migrate() error {
	goose.SetBaseFS(combinedMigrations)
	if !d.isMigrate {
		return nil
	}

	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Println("Error Migrating database: ", err)
		return err
	}

	if err := goose.Up(d.db, "migrations"); err != nil {
		fmt.Println("Error Migrating database: ", err)
		return err
	}
	return nil
}

// func (d *database) GetCartCoupons() []string {
// 	value := []string{"c1", "c2", "c3"}
// 	fmt.Println("GetCartCoupons DB: ", value)
// 	return value
// }
