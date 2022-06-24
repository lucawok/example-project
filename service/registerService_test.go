package service_test

import (
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

	data := &mongo.DeleteResult{DeletedCount: 1}

	fakeDB.DeleteByIDReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, err := serviceInstance.DeleteEmployeeById("2")
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, data, actual)
}
