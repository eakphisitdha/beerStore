package service

import (
	"beerstore/model"
	mocksRepository "beerstore/repository/mocks"
	mocksTransaction "beerstore/transaction/mocks"
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

func (suite *ServiceTestSuite) TestUpdate() {
	//Input
	req := model.UpdateRequest{
		ID:   1,
		Name: "A",
		Type: "can",
	}

	//expected
	setfields := []string{"name = 'A'", "type = 'can'"}
	fields := []string{"name", "type"}
	values := []interface{}{}
	values = append(values, "A", "can")
	suite.mockRepository.EXPECT().Update(req.ID, setfields).Return(nil)
	suite.mockTransaction.EXPECT().Log(fields, values, req.User, req.ID, "UPDATE").Return(nil)

	//function test
	err := suite.service.UpdateBeer(req)

	suite.Require().NoError(err)
}

func (suite *ServiceTestSuite) TestAdd() {
	//Input
	req := model.AddRequest{
		Name: "A",
		Type: "can",
		User: "user",
	}

	//expected
	fields := []string{"name", "type"}
	values := []interface{}{}
	values = append(values, "A", "can")
	suite.mockRepository.EXPECT().Add(fields, values).Return(1, nil)
	suite.mockTransaction.EXPECT().Log(fields, values, req.User, 1, "ADD").Return(nil)

	//function test
	err := suite.service.AddBeer(req)

	suite.Require().NoError(err)
}
