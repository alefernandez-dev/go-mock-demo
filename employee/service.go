package employee

import (
	"errors"
)

type EmployeeService struct {
	repository Employees
}

func NewEmployeeService(repository Employees) *EmployeeService {
	return &EmployeeService{repository: repository}
}

func (es EmployeeService) Register(employee Employee) error {

	if es.repository.ExistsById(employee.id) {
		return errors.New("employee already exists")
	}

	success := es.repository.Persist(employee)

	if !success {
		return errors.New("error to persist employee")
	}

	return nil
}

func (es EmployeeService) RetrieveAll(limit int16) []Employee {
	return es.repository.List(limit)
}

func (es EmployeeService) RetrieveById(employeeId int16) (*Employee, error) {

	employee, success := es.repository.GetById(employeeId)

	if !success {
		return &Employee{}, errors.New("employee not found")
	}

	return employee, nil

}
