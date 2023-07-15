package service

import (
	"beerstore/model"
	"beerstore/repository"
	"beerstore/transaction"
	"fmt"
	"log"
)

//go:generate mockery --with-expecter --name "IService" --output $PWD/mocks
type IService interface {
	GetBeers(req model.GetRequest) []model.BeerResponse
	AddBeer(req model.AddRequest) error
	UpdateBeer(req model.UpdateRequest) error
	DeleteBeer(id int, req model.DeleteRequest) error
}

type Service struct {
	r repository.IRepository
	t transaction.ITransaction
}

func NewService(r repository.IRepository, t transaction.ITransaction) IService {
	return &Service{r: r, t: t}
}

func (s *Service) GetBeers(req model.GetRequest) []model.BeerResponse {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 3
	}
	beers, err := s.r.Find(req)
	if err != nil {
		log.Println(err.Error())
	}
	var response []model.BeerResponse
	for _, beer := range beers {
		var res model.BeerResponse
		res.ID = beer.ID
		res.Name = beer.Name
		if beer.Type.Valid {
			res.Type = beer.Type.String
		}
		if beer.Detail.Valid {
			res.Detail = beer.Detail.String
		}
		if beer.URL.Valid {
			res.URL = beer.URL.String
		}
		response = append(response, res)
	}
	return response
}

func (s *Service) AddBeer(req model.AddRequest) error {
	fields := []string{}
	values := []interface{}{}

	// Check and add non-empty fields to query
	if req.Name != "" {
		fields = append(fields, "name")
		values = append(values, req.Name)
	}
	if req.Type != "" {
		fields = append(fields, "type")
		values = append(values, req.Type)
	}
	if req.Detail != "" {
		fields = append(fields, "detail")
		values = append(values, req.Detail)
	}
	if req.URL != "" {
		fields = append(fields, "url")
		values = append(values, req.URL)
	}

	id, err := s.r.Add(fields, values)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log(fields, values, req.User, id, "ADD")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) UpdateBeer(req model.UpdateRequest) error {
	var fields, setFields []string
	values := []interface{}{}

	// Check and add non-empty fields to query
	if req.Name != "" {
		setFields = append(setFields, fmt.Sprintf("name = '%s'", req.Name))
		fields = append(fields, "name")
		values = append(values, req.Name)
	}
	if req.Type != "" {
		setFields = append(setFields, fmt.Sprintf("type = '%s'", req.Type))
		fields = append(fields, "type")
		values = append(values, req.Type)
	}
	if req.Detail != "" {
		setFields = append(setFields, fmt.Sprintf("detail = '%s'", req.Detail))
		fields = append(fields, "detail")
		values = append(values, req.Detail)
	}
	if req.URL != "" {
		setFields = append(setFields, fmt.Sprintf("url = '%s'", req.URL))
		fields = append(fields, "url")
		values = append(values, req.URL)
	}

	err := s.r.Update(req.ID, setFields)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = s.t.Log(fields, values, req.User, req.ID, "UPDATE")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) DeleteBeer(id int, req model.DeleteRequest) error {

	err := s.r.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log(nil, nil, req.User, id, "DELETE")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
