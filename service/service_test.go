package service

import (
	"beerstore/model"
	mocksRepository "beerstore/repository/mocks"
	mocksTransaction "beerstore/transaction/mocks"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	mockRepository  *mocksRepository.IRepository
	mockTransaction *mocksTransaction.ITransaction
	service         IService
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.mockRepository = mocksRepository.NewIRepository(suite.T())
	suite.mockTransaction = mocksTransaction.NewITransaction(suite.T())
	suite.service = NewService(suite.mockRepository, suite.mockTransaction)
}

func (suite *ServiceTestSuite) TestGet() {
	//Input
	req := model.GetRequest{
		Name:     "LEO",
		Page:     1,
		PageSize: 10,
	}

	//mock
	suite.mockRepository.EXPECT().Find(req).Return([]model.Beer{
		{ID: 1,
			Name:   "LEO",
			Type:   sql.NullString{String: "Pale Lager", Valid: true},
			Detail: sql.NullString{String: "เบียร์แบรนด์ LEO", Valid: true},
			URL:    sql.NullString{String: "https://example.com/LEO.jpg", Valid: true}},
	}, nil)

	//expected
	expected := []model.BeerResponse{
		{ID: 1,
			Name:   "LEO",
			Type:   "Pale Lager",
			Detail: "เบียร์แบรนด์ LEO",
			URL:    "https://example.com/LEO.jpg"},
	}
	//function test
	actual, err := suite.service.GetBeers(req)

	suite.Require().NoError(err)
	suite.Equal(expected, actual)
}

func (suite *ServiceTestSuite) TestAdd() {
	//Input
	req := model.AddRequest{
		Name: "LEO",
		Type: "Pale Lager",
		User: "user",
	}

	//expected
	fields := []string{"name", "type"}
	values := []interface{}{}
	values = append(values, "LEO", "Pale Lager")
	suite.mockRepository.EXPECT().Add(fields, values).Return(1, nil)
	suite.mockTransaction.EXPECT().Log(fields, values, req.User, 1, "ADD").Return(nil)

	//function test
	err := suite.service.AddBeer(req)

	suite.Require().NoError(err)
}

func (suite *ServiceTestSuite) TestUpdate() {
	//Input
	req := model.UpdateRequest{
		ID:   1,
		Name: "LEO",
		Type: "Pale Lager",
	}

	//expected
	setfields := []string{"name = 'LEO'", "type = 'Pale Lager'"}
	fields := []string{"name", "type"}
	values := []interface{}{}
	values = append(values, "LEO", "Pale Lager")
	suite.mockRepository.EXPECT().Update(req.ID, setfields).Return(nil)
	suite.mockTransaction.EXPECT().Log(fields, values, req.User, req.ID, "UPDATE").Return(nil)

	//function test
	err := suite.service.UpdateBeer(req)

	suite.Require().NoError(err)
}
