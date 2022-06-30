// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"example-project/handler"
	"example-project/model"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type FakeServiceInterface struct {
	CreateEmployeesStub        func([]model.Employee) interface{}
	createEmployeesMutex       sync.RWMutex
	createEmployeesArgsForCall []struct {
		arg1 []model.Employee
	}
	createEmployeesReturns struct {
		result1 interface{}
	}
	createEmployeesReturnsOnCall map[int]struct {
		result1 interface{}
	}
	DeleteEmployeeByIdStub        func(string) (*mongo.DeleteResult, error)
	deleteEmployeeByIdMutex       sync.RWMutex
	deleteEmployeeByIdArgsForCall []struct {
		arg1 string
	}
	deleteEmployeeByIdReturns struct {
		result1 *mongo.DeleteResult
		result2 error
	}
	deleteEmployeeByIdReturnsOnCall map[int]struct {
		result1 *mongo.DeleteResult
		result2 error
	}
	GetAllEmployeesStub        func() ([]model.Employee, error)
	getAllEmployeesMutex       sync.RWMutex
	getAllEmployeesArgsForCall []struct {
	}
	getAllEmployeesReturns struct {
		result1 []model.Employee
		result2 error
	}
	getAllEmployeesReturnsOnCall map[int]struct {
		result1 []model.Employee
		result2 error
	}
	GetEmployeeByIdStub        func(string) model.Employee
	getEmployeeByIdMutex       sync.RWMutex
	getEmployeeByIdArgsForCall []struct {
		arg1 string
	}
	getEmployeeByIdReturns struct {
		result1 model.Employee
	}
	getEmployeeByIdReturnsOnCall map[int]struct {
		result1 model.Employee
	}
	GetPaginatedEmployeesStub        func(int, int) ([]model.Employee, error)
	getPaginatedEmployeesMutex       sync.RWMutex
	getPaginatedEmployeesArgsForCall []struct {
		arg1 int
		arg2 int
	}
	getPaginatedEmployeesReturns struct {
		result1 []model.Employee
		result2 error
	}
	getPaginatedEmployeesReturnsOnCall map[int]struct {
		result1 []model.Employee
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceInterface) CreateEmployees(arg1 []model.Employee) interface{} {
	var arg1Copy []model.Employee
	if arg1 != nil {
		arg1Copy = make([]model.Employee, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createEmployeesMutex.Lock()
	ret, specificReturn := fake.createEmployeesReturnsOnCall[len(fake.createEmployeesArgsForCall)]
	fake.createEmployeesArgsForCall = append(fake.createEmployeesArgsForCall, struct {
		arg1 []model.Employee
	}{arg1Copy})
	stub := fake.CreateEmployeesStub
	fakeReturns := fake.createEmployeesReturns
	fake.recordInvocation("CreateEmployees", []interface{}{arg1Copy})
	fake.createEmployeesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceInterface) CreateEmployeesCallCount() int {
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	return len(fake.createEmployeesArgsForCall)
}

func (fake *FakeServiceInterface) CreateEmployeesCalls(stub func([]model.Employee) interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = stub
}

func (fake *FakeServiceInterface) CreateEmployeesArgsForCall(i int) []model.Employee {
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	argsForCall := fake.createEmployeesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) CreateEmployeesReturns(result1 interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = nil
	fake.createEmployeesReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) CreateEmployeesReturnsOnCall(i int, result1 interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = nil
	if fake.createEmployeesReturnsOnCall == nil {
		fake.createEmployeesReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.createEmployeesReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) DeleteEmployeeById(arg1 string) (*mongo.DeleteResult, error) {
	fake.deleteEmployeeByIdMutex.Lock()
	ret, specificReturn := fake.deleteEmployeeByIdReturnsOnCall[len(fake.deleteEmployeeByIdArgsForCall)]
	fake.deleteEmployeeByIdArgsForCall = append(fake.deleteEmployeeByIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteEmployeeByIdStub
	fakeReturns := fake.deleteEmployeeByIdReturns
	fake.recordInvocation("DeleteEmployeeById", []interface{}{arg1})
	fake.deleteEmployeeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdCallCount() int {
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	return len(fake.deleteEmployeeByIdArgsForCall)
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdCalls(stub func(string) (*mongo.DeleteResult, error)) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = stub
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdArgsForCall(i int) string {
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	argsForCall := fake.deleteEmployeeByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdReturns(result1 *mongo.DeleteResult, result2 error) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = nil
	fake.deleteEmployeeByIdReturns = struct {
		result1 *mongo.DeleteResult
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdReturnsOnCall(i int, result1 *mongo.DeleteResult, result2 error) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = nil
	if fake.deleteEmployeeByIdReturnsOnCall == nil {
		fake.deleteEmployeeByIdReturnsOnCall = make(map[int]struct {
			result1 *mongo.DeleteResult
			result2 error
		})
	}
	fake.deleteEmployeeByIdReturnsOnCall[i] = struct {
		result1 *mongo.DeleteResult
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetAllEmployees() ([]model.Employee, error) {
	fake.getAllEmployeesMutex.Lock()
	ret, specificReturn := fake.getAllEmployeesReturnsOnCall[len(fake.getAllEmployeesArgsForCall)]
	fake.getAllEmployeesArgsForCall = append(fake.getAllEmployeesArgsForCall, struct {
	}{})
	stub := fake.GetAllEmployeesStub
	fakeReturns := fake.getAllEmployeesReturns
	fake.recordInvocation("GetAllEmployees", []interface{}{})
	fake.getAllEmployeesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetAllEmployeesCallCount() int {
	fake.getAllEmployeesMutex.RLock()
	defer fake.getAllEmployeesMutex.RUnlock()
	return len(fake.getAllEmployeesArgsForCall)
}

func (fake *FakeServiceInterface) GetAllEmployeesCalls(stub func() ([]model.Employee, error)) {
	fake.getAllEmployeesMutex.Lock()
	defer fake.getAllEmployeesMutex.Unlock()
	fake.GetAllEmployeesStub = stub
}

func (fake *FakeServiceInterface) GetAllEmployeesReturns(result1 []model.Employee, result2 error) {
	fake.getAllEmployeesMutex.Lock()
	defer fake.getAllEmployeesMutex.Unlock()
	fake.GetAllEmployeesStub = nil
	fake.getAllEmployeesReturns = struct {
		result1 []model.Employee
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetAllEmployeesReturnsOnCall(i int, result1 []model.Employee, result2 error) {
	fake.getAllEmployeesMutex.Lock()
	defer fake.getAllEmployeesMutex.Unlock()
	fake.GetAllEmployeesStub = nil
	if fake.getAllEmployeesReturnsOnCall == nil {
		fake.getAllEmployeesReturnsOnCall = make(map[int]struct {
			result1 []model.Employee
			result2 error
		})
	}
	fake.getAllEmployeesReturnsOnCall[i] = struct {
		result1 []model.Employee
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetEmployeeById(arg1 string) model.Employee {
	fake.getEmployeeByIdMutex.Lock()
	ret, specificReturn := fake.getEmployeeByIdReturnsOnCall[len(fake.getEmployeeByIdArgsForCall)]
	fake.getEmployeeByIdArgsForCall = append(fake.getEmployeeByIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetEmployeeByIdStub
	fakeReturns := fake.getEmployeeByIdReturns
	fake.recordInvocation("GetEmployeeById", []interface{}{arg1})
	fake.getEmployeeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceInterface) GetEmployeeByIdCallCount() int {
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	return len(fake.getEmployeeByIdArgsForCall)
}

func (fake *FakeServiceInterface) GetEmployeeByIdCalls(stub func(string) model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = stub
}

func (fake *FakeServiceInterface) GetEmployeeByIdArgsForCall(i int) string {
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	argsForCall := fake.getEmployeeByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) GetEmployeeByIdReturns(result1 model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = nil
	fake.getEmployeeByIdReturns = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeServiceInterface) GetEmployeeByIdReturnsOnCall(i int, result1 model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = nil
	if fake.getEmployeeByIdReturnsOnCall == nil {
		fake.getEmployeeByIdReturnsOnCall = make(map[int]struct {
			result1 model.Employee
		})
	}
	fake.getEmployeeByIdReturnsOnCall[i] = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeServiceInterface) GetPaginatedEmployees(arg1 int, arg2 int) ([]model.Employee, error) {
	fake.getPaginatedEmployeesMutex.Lock()
	ret, specificReturn := fake.getPaginatedEmployeesReturnsOnCall[len(fake.getPaginatedEmployeesArgsForCall)]
	fake.getPaginatedEmployeesArgsForCall = append(fake.getPaginatedEmployeesArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.GetPaginatedEmployeesStub
	fakeReturns := fake.getPaginatedEmployeesReturns
	fake.recordInvocation("GetPaginatedEmployees", []interface{}{arg1, arg2})
	fake.getPaginatedEmployeesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetPaginatedEmployeesCallCount() int {
	fake.getPaginatedEmployeesMutex.RLock()
	defer fake.getPaginatedEmployeesMutex.RUnlock()
	return len(fake.getPaginatedEmployeesArgsForCall)
}

func (fake *FakeServiceInterface) GetPaginatedEmployeesCalls(stub func(int, int) ([]model.Employee, error)) {
	fake.getPaginatedEmployeesMutex.Lock()
	defer fake.getPaginatedEmployeesMutex.Unlock()
	fake.GetPaginatedEmployeesStub = stub
}

func (fake *FakeServiceInterface) GetPaginatedEmployeesArgsForCall(i int) (int, int) {
	fake.getPaginatedEmployeesMutex.RLock()
	defer fake.getPaginatedEmployeesMutex.RUnlock()
	argsForCall := fake.getPaginatedEmployeesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceInterface) GetPaginatedEmployeesReturns(result1 []model.Employee, result2 error) {
	fake.getPaginatedEmployeesMutex.Lock()
	defer fake.getPaginatedEmployeesMutex.Unlock()
	fake.GetPaginatedEmployeesStub = nil
	fake.getPaginatedEmployeesReturns = struct {
		result1 []model.Employee
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetPaginatedEmployeesReturnsOnCall(i int, result1 []model.Employee, result2 error) {
	fake.getPaginatedEmployeesMutex.Lock()
	defer fake.getPaginatedEmployeesMutex.Unlock()
	fake.GetPaginatedEmployeesStub = nil
	if fake.getPaginatedEmployeesReturnsOnCall == nil {
		fake.getPaginatedEmployeesReturnsOnCall = make(map[int]struct {
			result1 []model.Employee
			result2 error
		})
	}
	fake.getPaginatedEmployeesReturnsOnCall[i] = struct {
		result1 []model.Employee
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	fake.getAllEmployeesMutex.RLock()
	defer fake.getAllEmployeesMutex.RUnlock()
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	fake.getPaginatedEmployeesMutex.RLock()
	defer fake.getPaginatedEmployeesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handler.ServiceInterface = new(FakeServiceInterface)
