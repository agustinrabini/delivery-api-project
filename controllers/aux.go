package controllers

import (
	"delivery-api-project/controllers/web/request"
	"delivery-api-project/controllers/web/response"
	"fmt"
)

func validateOrder(order response.Order) error {

	if order.Id == 0 || order.ReceiverID == 0 || order.Status == "" || order.CreationDate == "" {
		return fmt.Errorf("error receiving nil values for order, expect non nil values")
	}

	if order.Delivery.OriginLocation.Type == "" || order.Delivery.OriginLocation.Province == "" || order.Delivery.OriginLocation.City == "" || order.Delivery.OriginLocation.Commune == "" || order.Delivery.OriginLocation.FullAddress == "" || order.Delivery.OriginLocation.Lat == 0 || order.Delivery.OriginLocation.Lng == 0 {
		return fmt.Errorf("error receiving nil values for deliveries locations, expect non nil values")
	}

	for _, pkg := range order.Packages {

		if pkg.Id == 0 || pkg.IdOrder == 0 || pkg.QuantityItems == 0 || pkg.Size == "" || pkg.Weight == 0 {
			return fmt.Errorf("error receiving nil values for packages, expect non nil values")
		}
	}

	return nil
}

func validateRequest(order request.Order) error {

	for _, pkg := range order.Packages {

		if pkg.QuantityItems < 1 {
			return fmt.Errorf("the quantity of items cannot be less than one")
		}

		if pkg.Weight > 25 {
			return fmt.Errorf("the service for packages that weights more than 25kg is not avaible, order not taked")
		}
		if pkg.Weight < 0 {
			return fmt.Errorf("please enter the weigth of the package")
		}
	}

	if order.Delivery.DestinyLocation.Province == "" || order.Delivery.OriginLocation.Type == "" || order.Delivery.OriginLocation.Commune == "" || order.Delivery.DestinyLocation.FullAddress == "" || order.Delivery.DestinyLocation.Lng == 0 || order.Delivery.DestinyLocation.Lat == 0 {
		return fmt.Errorf("missing values on the request, please check the body")
	}
	return nil

}
