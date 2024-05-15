package employee_test

import (
	"testing"

	"github.com/alefernandez-dev/go-mock-demo/employee"
)

type EmployeeRepositoryMock struct {
	employee  employee.Employee
	success   bool
	exits     bool
	employees []employee.Employee
}

func (erm EmployeeRepositoryMock) Persist(employee employee.Employee) bool {
	return erm.success
}
func (erm EmployeeRepositoryMock) ExistsById(employeeId int16) bool {
	return erm.exits
}
func (erm EmployeeRepositoryMock) GetById(employeeId int16) (*employee.Employee, bool) {
	return &erm.employee, erm.success
}
func (erm EmployeeRepositoryMock) List(limit int16) []employee.Employee {
	return erm.employees
}

func TestRegisterNewEmployee(t *testing.T) {

	employeeRepositoryMock := &EmployeeRepositoryMock{}

	service := employee.NewEmployeeService(employeeRepositoryMock)

	t.Run("register a valid employee", func(t *testing.T) {

		employeeRepositoryMock.exits = false
		employeeRepositoryMock.success = true

		err := service.Register(*employee.NewEmployee(1, "Alejandro", "Fernandez", 120.500, true))
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

	})

	t.Run("register an employee with duplicate ID", func(t *testing.T) {

		employeeRepositoryMock.exits = true
		employeeRepositoryMock.success = true

		err := service.Register(*employee.NewEmployee(1, "Alejandro", "Fernandez", 120.500, true))
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}

	})

	t.Run("retrieve emlpoyee by ID successfully", func(t *testing.T) {

		employeeRepositoryMock.employee = *employee.NewEmployee(1, "Alejandro", "Fernandez", 120.500, true)
		employeeRepositoryMock.success = true

		employee, err := service.RetrieveById(1)

		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

		if got := employee.Firstname(); got != "Alejandro" {
			t.Fatalf("expected 'Alejandro', got %q", got)
		}

	})

	t.Run("failed to retrieve employee by ID", func(t *testing.T) {

		employeeRepositoryMock.success = false

		_, err := service.RetrieveById(1)

		if err == nil {
			t.Errorf("expected error, got %v", err)
		}

	})

}
