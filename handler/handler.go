package handler

import (
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) interface{}
	GetEmployeeById(id string) model.Employee
	GetAllEmployees() ([]model.Employee, error)
	DeleteEmployeeById(id string) (*mongo.DeleteResult, error)
	GetPaginatedEmployees(page int, limit int) (model.PaginatedPayload, error)
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
		var employees []model.Employee
		employees = append(employees, employee)
		userErrorCheck, _ := handler.DoUserExist(employees)
		if userErrorCheck == true {
			c.AbortWithStatusJSON(400, gin.H{
				"errorMessage": "User already exists with this ID",
			})
			return
		}

		response := handler.ServiceInterface.CreateEmployees(employees)
		c.JSON(200, response)
		return
	} else {
		userErrorCheck, ErrorArray := handler.DoUserExist(payLoad.Employees)
		if userErrorCheck == true {
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

func (handler Handler) DoUserExist(emp []model.Employee) (bool, []model.Employee) {
	var idList []string
	var errorEmployees []model.Employee

	for _, employee := range emp {
		response := handler.ServiceInterface.GetEmployeeById(employee.ID)
		if len(response.ID) != 0 {
			errorEmployees = append(errorEmployees, employee)
		} else {
			idList = append(idList, employee.ID)
			var idCount int = 0
			for _, id := range idList {
				if id == employee.ID {
					idCount++
				}
			}
			if idCount >= 2 {
				errorEmployees = append(errorEmployees, employee)
			}
		}
	}
	if len(errorEmployees) != 0 {
		return true, errorEmployees
	} else {
		return false, nil
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
	pages, pageOk := c.GetQuery("page")
	limit, limitOk := c.GetQuery("limit")
	pageInt, pageErr := strconv.Atoi(pages)
	limitInt, limitErr := strconv.Atoi(limit)
	if pageOk && limitOk {
		if pageOk && limitOk && pageErr == nil && limitErr == nil {

			response, err := handler.ServiceInterface.GetPaginatedEmployees(pageInt, limitInt)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"errorMessage": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, response)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": "queries are invalid, please check or remove them",
			})
			return
		}
	} else {
		response, err := handler.ServiceInterface.GetAllEmployees()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}

}
