package service_test

import (
	"errors"
	"example-project/model"
	"example-project/service"
	"example-project/service/servicefakes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}

	fakeDB.GetByIDReturns(data)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.GetEmployeeById("1")
	assert.Equal(t, data, actual)

}

func TestDeleteEmployeeById(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	noUserError := errors.New("no user deleted, please check the id")
	data := &mongo.DeleteResult{DeletedCount: 0}
	fakeDB.DeleteByIDReturns(&mongo.DeleteResult{DeletedCount: 0}, noUserError)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, err := serviceInstance.DeleteEmployeeById("3")
	assert.Equal(t, data, actual, err)
}

func TestGetAllEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	var dataArray []model.Employee
	var data model.Employee
	data.ID = "1"
	data.Email = "example@gmail.com"
	data.FirstName = "Test"
	data.LastName = "Tester"
	dataArray = append(dataArray, data)
	fakeDB.GetAllReturns(dataArray, nil)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, err := serviceInstance.GetAllEmployees()
	assert.Equal(t, dataArray, actual, err)
}
