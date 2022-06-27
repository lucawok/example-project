package handler_test

import (
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
