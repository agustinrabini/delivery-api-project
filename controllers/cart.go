package controllers

import (
	"strconv"

	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"

	"github.com/agustinrabini/jopit-api-carts/src/api/controllers/web/request"
	response "github.com/agustinrabini/jopit-api-carts/src/api/controllers/web/response"

	"github.com/agustinrabini/jopit-api-carts/src/api/internal/cart"

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

type CartController struct {
	s cart.Service
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

func NewCartController(s cart.Service) *CartController {
	return &CartController{
		s: s}
}

//Retrieves the active cart of the shop for a determinated user.
func (cn *CartController) GetActive() gin.HandlerFunc {
	return func(c *gin.Context) {

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		finalCart, integrity, aerr := cn.s.GetActive(c, uId, &idShop)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		if !integrity {
			c.Header("warning", "ITEM_UNAVAILABLE")
		}

		c.JSON(200, finalCart)
	}
}

func (cn *CartController) GetItemsCount() gin.HandlerFunc {
	return func(c *gin.Context) {

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		numberItems, aerr := cn.s.GetItemsCount(c, uId, &idShop)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		c.JSON(200, *numberItems)
	}
}

func (cn *CartController) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		var i request.AddItemRequest

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		err = c.ShouldBindJSON(&i)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageCreate, "Unprocessable Entity", 422, apierrors.CauseList{"Check the body of the request"})
			c.JSON(422, apiErr)
			return
		}

		integrity, aerr := cn.s.AddItem(c, i, uId, &idShop)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		if !integrity {
			c.Header("warning", "ITEM_UNAVAILABLE")
		}

		c.JSON(204, "")
	}
}

func (cn *CartController) CartConfirm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var i []request.CartItem

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		err = c.ShouldBindJSON(&i)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageCreate, "Unprocessable Entity", 422, apierrors.CauseList{"Check the body of the request"})
			c.JSON(422, apiErr)
			return
		}

		integrity, aerr := cn.s.CartConfirm(c, uId, &idShop, i)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		if !integrity {
			c.Header("warning", "ITEM_UNAVAILABLE")
			c.JSON(204, "one of the items is not avaible.")
		}

		c.JSON(204, "")
	}
}

func (cn *CartController) GetActiveCartId() gin.HandlerFunc {
	return func(c *gin.Context) {

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		finalCart, aerr := cn.s.GetActiveCartId(c, uId, &idShop)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		c.JSON(200, finalCart)
	}
}

func (cn *CartController) GetActiveCartIdAndPrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		uId := c.GetString("user_id")

		sId := c.Param("id_shop")
		idShop, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(400, apierrors.NewApiError("check the id in the path.", err.Error(), 400, apierrors.CauseList{}))
			return
		}

		idCart, total, aerr := cn.s.GetActiveCartIdAndPrice(c, uId, &idShop)
		if aerr != nil {
			c.JSON(aerr.Status(), aerr)
			return
		}

		cart := response.CartIdAndPrice{
			Id:    *idCart,
			Price: *total,
		}

		c.JSON(200, cart)
	}
}
