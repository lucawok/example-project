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

	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeDB.GetByIDReturns(data)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.GetEmployeeById("1")
	assert.Equal(t, data, actual)

}

func TestDeleteEmployeeById(t *testing.T) {
	data := &mongo.DeleteResult{DeletedCount: 0}
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeDB.DeleteByIDReturns(data, nil)
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

func TestCreateEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	var fakeEmployees []model.Employee
	emp1 := model.Employee{ID: "1", FirstName: "Test", LastName: "Tester", Email: "example@gmail.com"}
	emp2 := model.Employee{ID: "2", FirstName: "Thomas", LastName: "Crock", Email: "example@gmx.de"}
	fakeEmployees = append(fakeEmployees, emp1)
	fakeEmployees = append(fakeEmployees, emp2)
	fakeDB.UpdateManyReturns(fakeEmployees)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.CreateEmployees(fakeEmployees)
	assert.Equal(t, fakeEmployees, actual)
}

func TestCreateOneEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	var fakeEmployees []model.Employee
	emp1 := model.Employee{ID: "1", FirstName: "Test", LastName: "Tester", Email: "example@gmail.com"}
	fakeEmployees = append(fakeEmployees, emp1)
	fakeDB.UpdateOneReturns(fakeEmployees)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.CreateEmployees(fakeEmployees)
	assert.Equal(t, fakeEmployees, actual)
}
