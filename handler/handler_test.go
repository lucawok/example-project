package handler_test

import (
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "d", Value: "e32rr2r23r"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeByIdReturns("1")

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteEmployeeHandler(fakeContext)

	assert.Equal(t, 400, responseRecorder.Code)
}

func TestDeleteEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeByIdReturns("1")

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteEmployeeHandler(fakeContext)

	assert.Equal(t, 200, responseRecorder.Code)
}
