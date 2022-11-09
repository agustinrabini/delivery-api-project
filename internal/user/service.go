package user

import (
	"errors"
	"fmt"

	"github.com/agustinrabini/poc-proximity-objects/calculus"
	"github.com/agustinrabini/poc-proximity-objects/domain"
	"github.com/agustinrabini/poc-proximity-objects/internal/commune"
	"github.com/agustinrabini/poc-proximity-objects/internal/communesshop"
	"github.com/agustinrabini/poc-proximity-objects/internal/shop"
)

type Service interface {
	GetOne(id int) (domain.User, error)
	GetClosestShops(user_id int, refFilter float64, usrFilter float64) (domain.User, domain.Shops, error)
}

type service struct {
	uRespository Repository
	cRepository  commune.Repository
	sRepository  shop.Repository
	csRepository communesshop.Repository
}

func NewService(uRepository Repository, cRepository commune.Repository, sRepository shop.Repository, csRepository communesshop.Repository) Service {
	return &service{
		uRespository: uRepository,
		cRepository:  cRepository,
		sRepository:  sRepository,
		csRepository: csRepository,
	}
}

func (s *service) GetOne(id int) (domain.User, error) {

	usr, err := s.uRespository.GetUser(id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	return usr, nil
}

func (s *service) GetClosestShops(user_id int, refFilter float64, usrFilter float64) (domain.User, domain.Shops, error) {

	var usr domain.User
	var shops domain.Shops
	var shop domain.Shop
	var cms domain.Communes
	var cmshs domain.CommunesShops
	var clossestRefPointsID []int
	var err error

	usr, err = s.uRespository.GetUser(user_id)
	if err != nil {
		return domain.User{}, domain.Shops{}, errors.New(err.Error())
	}

	cms, err = s.cRepository.GetAll()
	if err != nil {
		return domain.User{}, domain.Shops{}, errors.New(err.Error())
	}

	//calculates the distance between an user and the reference points, return an slice of int with the commune number
	//which works as the ID.
	for _, c := range cms {

		refDist := calculus.Distance(usr.Lat, usr.Lng, c.Lat, c.Lng)

		if refDist < refFilter {
			clossestRefPointsID = append(clossestRefPointsID, c.Number)
		}
	}

	//brings shops for the given communes
	for _, c := range clossestRefPointsID {

		cmshs, err = s.csRepository.GetShopByCommune(c)
		if err != nil {
			return domain.User{}, domain.Shops{}, errors.New(err.Error())
		}

		for _, cmsh := range cmshs {

			shop, err = s.sRepository.GetOne(cmsh.Shop_id)
			if err != nil {
				return domain.User{}, domain.Shops{}, errors.New(err.Error())
			}

			distUsrToShop := calculus.Distance(usr.Lat, usr.Lng, shop.Lat, shop.Lng)
			if distUsrToShop < usrFilter {
				shops = append(shops, shop)
			}
		}

	}

	if len(shops) == 0 {
		return domain.User{}, domain.Shops{}, fmt.Errorf("No shops founded under thsi radius: %f", usrFilter)
	}

	return usr, shops, nil
}
