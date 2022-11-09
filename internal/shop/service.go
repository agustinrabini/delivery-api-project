package shop

import (
	"errors"

	"github.com/agustinrabini/delivery-api-project/domain"
)

type Service interface {
	GetAll() (domain.Shops, error)
	GetOne(id int) (domain.Shop, error)
}

type service struct {
	respository Repository
}

func NewService(repository Repository) Service {
	return &service{respository: repository}
}

func (s *service) GetAll() (domain.Shops, error) {

	sh, err := s.respository.GetAll()
	if err != nil {
		return domain.Shops{}, errors.New(err.Error())
	}

	return sh, nil
}

func (s *service) GetOne(id int) (domain.Shop, error) {

	sh, err := s.respository.GetOne(id)
	if err != nil {
		return domain.Shop{}, errors.New(err.Error())
	}

	return sh, nil
}
