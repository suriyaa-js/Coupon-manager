package database

import "embed"

// //go:embed migrations/*_coupon_manage_table.sql
// var embedMigrateCoupon embed.FS

// //go:embed migrations/*_coupon_manage_table.sql
// var embedMigrateCouponCondition embed.FS

//go:embed migrations/*_new_coupon_manager.sql
var combinedMigrations embed.FS
