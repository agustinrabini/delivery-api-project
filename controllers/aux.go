package controllers

import (
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
