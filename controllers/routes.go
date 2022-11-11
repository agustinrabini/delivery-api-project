package controllers

import (
	"database/sql"
	"delivery-api-project/internal/delivery"
	"delivery-api-project/internal/location"
	"delivery-api-project/internal/order"
	"delivery-api-project/internal/packages"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {

	orderRepo := order.NewRepository(r.db)
	pkgsRepo := packages.NewRepository(r.db)
	deliveryRepo := delivery.NewRepository(r.db)
	locationRepo := location.NewRepository(r.db)

	service := order.NewService(orderRepo, pkgsRepo, deliveryRepo, locationRepo)
	handler := NewOrderController(service)

	r.rg.POST("/order/create", handler.Create())
	r.rg.GET("/order/update/:id/:status", handler.UpdateStatus())
	r.rg.GET("/sellers/get/:id", handler.Get())
}
