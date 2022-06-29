package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
	})

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContest)

	assert.Equal(t, http.StatusOK, responseRecoder.Code)

}

func TestGetEmployeeHandler_Return_invalid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "d", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
	})

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContest)

	assert.Equal(t, 400, responseRecoder.Code)

}

func TestDeleteEmployeeHandler_Return_invalid_status_code(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "123"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	errorNoUser := errors.New("no user deleted, please check the id")
	fakeService.DeleteEmployeeByIdReturns(&mongo.DeleteResult{DeletedCount: 0}, errorNoUser)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestDeleteEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeByIdReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteEmployeeHandler(fakeContext)

	assert.Equal(t, 200, responseRecorder.Code)
}

func TestDeleteEmployeeHandler_Return_invalid_params(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "i", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeByIdReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestGetAllEmployeesHandler_Return_valid_status_code(t *testing.T) {
	testEmployee := model.Employee{
		ID:        "1",
		Email:     "example@gmail.com",
		FirstName: "Test",
		LastName:  "Tester",
	}
	var testEmployeeArray []model.Employee
	testEmployeeArray = append(testEmployeeArray, testEmployee)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetAllEmployeesReturns(testEmployeeArray, nil)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, 200, responseRecorder.Code)
}

func TestGetAllEmployeesHandler_Return_invalid_status_code(t *testing.T) {
	databaseError := errors.New("no employees exist")

	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetAllEmployeesReturns(nil, databaseError)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, 400, responseRecorder.Code)
}

func TestCreateEmployeesHandler_SingleEmployee_valid_status_code(t *testing.T) {

	var jsonPayload = `
		{
		"id": "100",
		"first_name": "Test",
		"last_name": "Tester",
		"email": "tester@example.com"
		}
	`
	var employeePayload model.Employee
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.CreateEmployeeHandler(fakeContext)
	assert.Equal(t, 200, responseRecorder.Code)
}

func TestCreateEmployeesHandler_SingleEmployee_invalid_status_code(t *testing.T) {

	var jsonPayload = `
		{
		"id": "100",
		"first_name": "Test",
		"last_name: "Tester",
		"email": "tester@example.com"
		}
	`
	var employeePayload model.Employee
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.CreateEmployeeHandler(fakeContext)
	assert.Equal(t, 400, responseRecorder.Code)
}

func TestCreateEmployeesHandler_SingleEmployee_user_exists(t *testing.T) {

	var jsonPayload = `
		{
		"id": "100",
		"first_name": "Test",
		"last_name": "Tester",
		"email": "tester@example.com"
		}
	`
	var employeePayload model.Payload
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	fakeService.GetEmployeeByIdReturns(model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@example.com"})

	handlerInstance.CreateEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestCreateEmployeesHandler_MultiEmployee_valid_status_code(t *testing.T) {

	var jsonPayload = `
		{
		"employees": [
			{
			"id": "100",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			},
			{
			"id": "200",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			}
		]	
	}
	`
	var employeePayload model.Payload
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.CreateEmployeeHandler(fakeContext)
	assert.Equal(t, 200, responseRecorder.Code)
}

func TestCreateEmployeesHandler_MultiEmployee_invalid_data(t *testing.T) {

	var jsonPayload = `
		{
		"employees": [
			{
			"id": "100",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			},
			{
			"id": "200",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			
		]	
	}
	`
	var employeePayload model.Payload
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.CreateEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestCreateEmployeesHandler_MultiEmployee_user_exists(t *testing.T) {

	var jsonPayload = `
		{
		"employees": [
			{
			"id": "100",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			},
			{
			"id": "200",
			"first_name": "Test",
			"last_name": "Tester",
			"email": "tester@example.com"
			}
		]	
	}
	`
	var employeePayload model.Payload
	json.Unmarshal([]byte(jsonPayload), &employeePayload)
	body := bytes.NewBufferString(jsonPayload)
	responseRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", body)
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.CreateEmployeesReturns(fakeEmployees)
	responseRecorder.Body = body

	handlerInstance := handler.NewHandler(fakeService)
	fakeService.GetEmployeeByIdReturns(model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@example.com"})
	handlerInstance.CreateEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestDoUserExist_true_already_exists(t *testing.T) {
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	fakeService.GetEmployeeByIdReturns(fakeEmployees[0])
	handlerInstance := handler.NewHandler(fakeService)
	boolResult, _ := handlerInstance.DoUserExist(fakeEmployees)
	assert.Equal(t, true, boolResult)
}

func TestDoUserExist_true_duplication(t *testing.T) {
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}
	var emptyEmployee model.Employee
	fakeService.GetEmployeeByIdReturns(emptyEmployee)
	handlerInstance := handler.NewHandler(fakeService)
	boolResult, _ := handlerInstance.DoUserExist(fakeEmployees)
	assert.Equal(t, true, boolResult)
}

func TestDoUserExist_false(t *testing.T) {
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{},
	}
	fakeService.GetEmployeeByIdReturns(fakeEmployees[0])
	handlerInstance := handler.NewHandler(fakeService)
	boolResult, _ := handlerInstance.DoUserExist(fakeEmployees)
	assert.Equal(t, false, boolResult)
}
