package handler

import (
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) interface{}
	GetEmployeeById(id string) model.Employee
	GetAllEmployees() ([]model.Employee, error)
	DeleteEmployeeById(id string) (*mongo.DeleteResult, error)
}

type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(serviceInterface ServiceInterface) Handler {
	return Handler{
		ServiceInterface: serviceInterface,
	}
}

func (handler Handler) CreateEmployeeHandler(c *gin.Context) {
	var payLoad model.Payload
	err := c.ShouldBindBodyWith(&payLoad, binding.JSON)
	if err != nil || len(payLoad.Employees) == 0 {

		var employee model.Employee
		err := c.ShouldBindBodyWith(&employee, binding.JSON)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"errorMessage": "Invalid Data",
			})
			return
		}
		if handler.DoUserExist(employee) == true {
			c.AbortWithStatusJSON(400, gin.H{
				"errorMessage": "User already exists with this ID",
			})
			return
		}
		var employees []model.Employee
		employees = append(employees, employee)
		response := handler.ServiceInterface.CreateEmployees(employees)
		c.JSON(200, response)
		return
	} else {
		var ErrorArray []model.Employee
		for _, emp := range payLoad.Employees {
			if handler.DoUserExist(emp) {
				ErrorArray = append(ErrorArray, emp)
			}
		}
		if len(ErrorArray) != 0 {
			c.AbortWithStatusJSON(400, gin.H{
				"errorMessage": "The following employees need another ID",
				"Employees":    ErrorArray,
			})
			return
		}
		response := handler.ServiceInterface.CreateEmployees(payLoad.Employees)
		c.JSON(200, response)
		return
	}

}

func (handler Handler) DoUserExist(emp model.Employee) bool {
	response := handler.ServiceInterface.GetEmployeeById(emp.ID)
	if len(response.ID) == 0 {
		return false
	} else {
		return true
	}
}

func (handler Handler) GetEmployeeHandler(c *gin.Context) {
	pathParam, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "id is not given",
		})
		return
	}

	response := handler.ServiceInterface.GetEmployeeById(pathParam)
	c.JSON(http.StatusOK, response)
}

func (handler Handler) DeleteEmployeeHandler(c *gin.Context) {
	pathParam, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "ID is not given",
		})
		return
	}

	response, err := handler.ServiceInterface.DeleteEmployeeById(pathParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler Handler) GetAllEmployeesHandler(c *gin.Context) {
	response, err := handler.ServiceInterface.GetAllEmployees()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
