package controllers

import (
	"delivery-api-project/controllers/web/request"
	_ "delivery-api-project/docs"
	"delivery-api-project/internal/order"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	s order.Service
}

// Ping is the handler of test app
// @Summary Ping
// @Description test if the router works correctly
// @Tags ping
// @Produce  json
// @Param id path string true "Shop ID"
// @Success 200
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func NewOrderController(s order.Service) *OrderController {
	return &OrderController{
		s: s}
}

// @Summary Get order
// @Description Get order of an user
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} request.Order
// @Router /order/get/:id [get]
func (cn *OrderController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		order, err := cn.s.Get(c, id)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		if (order.Status == "") && err == nil {
			msg := fmt.Sprintf("any order found with with id: %d", id)
			c.JSON(200, msg)
		}

		err = validateOrder(order)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, order)
	}
}

func (cn *OrderController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		request := request.Order{}

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		err = validateRequest(request)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		id, err := cn.s.Create(c, request)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		msg := fmt.Sprintf("order created with id: %d", *id)

		c.JSON(200, msg)
	}
}

func (cn *OrderController) UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(402, err.Error())
			return
		}

		status := c.Param("status")

		err = cn.s.UpdateStatus(c, &id, status)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		c.Status(200)
	}
}
