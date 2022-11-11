package controllers

import (
	"delivery-api-project/controllers/web/request"
	"delivery-api-project/internal/order"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	genericErrorMessageGet       = "Error obtaining object"
	genericErrorMessageCreate    = "Error creating object"
	genericErrorMessageUpdate    = "Error updating object"
	genericErrorMessageDelete    = "Error deleting object"
	genericErrorMessageInvalidId = "Error invalid ID"
	genericErrorMessageNotFound  = "Resource not found"
)

type OrderController struct {
	s order.Service
}

// Ping is the handler of test app
// @Summary Ping
// @Description test if the router works correctly
// @Tags ping
// @Produce  json
// @Success 200
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func NewOrderController(s order.Service) *OrderController {
	return &OrderController{
		s: s}
}

func (cn *OrderController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(402, err.Error())
			return
		}

		order, err := cn.s.Get(c, id)
		if err != nil {
			c.JSON(500, err.Error())
			return
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

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		id, err := cn.s.Create(c, request)
		if err != nil {
			c.JSON(422, err.Error())
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
			c.JSON(402, err.Error())
			return
		}

		c.JSON(201, "")

	}
}
