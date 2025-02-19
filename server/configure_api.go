package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"suriyaa.com/coupon-manager/database"
	"suriyaa.com/coupon-manager/handlers"
	"suriyaa.com/coupon-manager/server/props"
	sc "suriyaa.com/coupon-manager/service"
)

func configureAPI(config *props.Config, s *server) {
	controller, err := initailizeHanlders(config)
	if err != nil {
		log.Fatalf("Error creating API controller in server: %s", err)
	}
	configureRoutes(controller, s)
}

func configureRoutes(api handlers.API, s *server) {
	// Define your routes here
	fmt.Println("Configuring routes")
	s.router.GET("/coupons", WrapHandler(api.GetCoupons))
	s.router.GET("/coupon/:id", WrapHandler(api.GetCouponByID))
	s.router.POST("/coupon", WrapHandler(api.AddCoupon))
	s.router.PUT("/coupon/:id", WrapHandler(api.UpdateCoupon))
	s.router.DELETE("/coupon/:id", WrapHandler(api.DeleteCoupon))

	// All Product routes
	s.router.POST("/product", WrapHandler(api.AddProduct))
	s.router.DELETE("/product/:id", WrapHandler(api.DeleteProduct))

	// Cart routes
	// s.router.POST("/cart", WrapHandler(api.AddItemToCart))

	s.router.GET("/", WrapHandler(api.ServeHTTP))
}

// initailizeHanlders creates a new instance of the handlers with couponService,
// initialize DB and CouponService
func initailizeHanlders(config *props.Config) (handlers.API, error) {
	dbConn, err := database.ConnectDB(config.Database)
	if err != nil {
		return nil, logError(err)
	}
	db := database.NewDatabase(dbConn, config.Database.IsMigrate)
	db.Migrate()
	couponService := sc.NewCouponService(db)
	controller := handlers.NewHandler(couponService)
	return controller, nil
}

// WrapHandler converts http.HandlerFunc to gin.HandlerFunc
func WrapHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id != "" {
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "id", id))
		}
		h(c.Writer, c.Request)
	}
}

func logError(err error) error {
	fmt.Println("Error in configure_api: ", err)
	return err
}
