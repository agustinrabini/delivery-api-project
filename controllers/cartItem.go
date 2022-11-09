package controllers

import (
	ci "github.com/agustinrabini/jopit-api-carts/src/api/internal/cartItem"
)

type CartItemController struct {
	s ci.Service
}

func NewCartItemController(s ci.Service) *CartItemController {
	return &CartItemController{
		s: s,
	}
}

/*
func (cn *CartItemController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ds, err := cn.s.GetAll(c)
		if err != nil {
			apiErr := apierrors.NewInternalServerApiError(genericErrorMessageGet, err)
			c.JSON(500, apiErr)
			return
		}
		if len(ds) == 0 {
			c.JSON(204, domain.Carts{})
			return
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, ds)
	}
}

func (cn *CartItemController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			apiErr := apierrors.NewBadRequestApiError(genericErrorMessageUpdate + " check the ID on the path")
			c.JSON(400, apiErr)
			return
		}
		d, err := cn.s.Get(c, &id)
		if err != nil {
			c.JSON(204, domain.Carts{})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(200, d)
	}
}

 func (cn *CartItemController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		p := domain.Cart{}

		err := c.ShouldBindJSON(&p)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageCreate, "Unprocessable Entity", 422, apierrors.CauseList{"Check the body of the request"})
			c.JSON(422, apiErr)
			return
		}

		err = validateCart(p)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageCreate, "Unprocessable Entity", 422, apierrors.CauseList{err.Error()})
			c.JSON(422, apiErr)
			return
		}

		err = validateItems(p.Items)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageCreate, "Unprocessable Entity", 422, apierrors.CauseList{err.Error()})
			c.JSON(422, apiErr)
			return
		}

		err = cn.s.Create(c, p)
		if err != nil {
			apiErr := apierrors.NewInternalServerApiError(genericErrorMessageCreate, err)
			c.JSON(500, apiErr)
			return
		}

		c.Status(201)
	}
}

func (cn *CartItemController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		p := domain.Cart{}

		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			apiErr := apierrors.NewBadRequestApiError(genericErrorMessageUpdate + " check the ID on the path")
			c.JSON(400, apiErr)
			return
		}

		err = c.ShouldBindJSON(&p)
		if err != nil {
			apiErr := apierrors.NewApiError(genericErrorMessageUpdate, "Unprocessable Entity", 422, apierrors.CauseList{"Check the body of the request"})
			c.JSON(422, apiErr)
			return
		}

		err = cn.s.Update(c, &id, p)
		if err != nil {
			c.JSON(204, "")
			return
		}

		c.Status(200)
	}
}

func (cn *CartItemController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		i := c.Param("id")
		id, err := strconv.Atoi(i)

		if err != nil {
			apiErr := apierrors.NewBadRequestApiError(genericErrorMessageUpdate + " check the ID on the path")
			c.JSON(400, apiErr)
			return
		}
		err = cn.s.Delete(c, &id)
		if err != nil {
			c.JSON(204, "")
			return
		}

		c.Status(200)
	}
} */

/* func (cn *CartItemController) GetItemsCount() gin.HandlerFunc {
	return func(c *gin.Context) {

		itemsCount, err := getItemsCount(c, cn)
		if err != nil {
			c.JSON(err.Status(), err)
			return
		}

		c.JSON(200, itemsCount)

	}
}

//It contains the logic to the GetActive() method.
func getItemsCount(c *gin.Context, cn *CartItemController) (*int, apierrors.ApiError) {

	//ok//
	idCart, err := cn.s.GetItemsCount(c, &idUser, &idShop)
	if err != nil {
		apiErr := apierrors.NewApiError("error from the server side", err.Error(), 500, apierrors.CauseList{})
		return nil, apiErr
	}

	return idCart, nil
}
*/
