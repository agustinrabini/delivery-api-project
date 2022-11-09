package commune

import (
	"errors"

	"github.com/agustinrabini/delivery-api-project/domain"
)

type Service interface {
	GetAll() (domain.Communes, error)
	GetOne(number int) (domain.Commune, error)
}

type service struct {
	respository Repository
}

func NewService(repository Repository) Service {
	return &service{respository: repository}
}

func (s *service) GetAll() (domain.Communes, error) {

	c, err := s.respository.GetAll()
	if err != nil {
		return domain.Communes{}, errors.New(err.Error())
	}

	return c, nil
}

func (s *service) GetOne(number int) (domain.Commune, error) {

	c, err := s.respository.GetOne(number)
	if err != nil {
		return domain.Commune{}, errors.New(err.Error())
	}

	return c, nil
}
