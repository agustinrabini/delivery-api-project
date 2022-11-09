package main

import (
	"fmt"
	"log"

	"github.com/agustinrabini/poc-proximity-objects/db"
	"github.com/agustinrabini/poc-proximity-objects/internal/commune"
	"github.com/agustinrabini/poc-proximity-objects/internal/communesshop"
	"github.com/agustinrabini/poc-proximity-objects/internal/shop"
	usr "github.com/agustinrabini/poc-proximity-objects/internal/user"
)

func main() {

	db := db.DbConect()

	userRepo := usr.NewRepository(db)
	communeRepo := commune.NewRepository(db)
	communesshopRepo := communesshop.NewRepository(db)
	shopRepo := shop.NewRepository(db)

	service := usr.NewService(userRepo, communeRepo, shopRepo, communesshopRepo)

	usrFilter := 2500.00
	refFilter := 10000.00
	usrID := 3

	usr, shops, err := service.GetClosestShops(usrID, refFilter, usrFilter)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, s := range shops {
		fmt.Printf("The shop %d is under %fmts for user %d \n\n", s.Shop_id, usrFilter, usr.ID)
	}
}
