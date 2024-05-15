package main

import (
	"fmt"

	"github.com/alefernandez-dev/go-mock-demo/employee"
)

func main() {

	employee01 := employee.NewEmployee(1, "Alejandro", "Fernández", 120000, true)
	employee02 := employee.NewEmployee(2, "María", "González", 90000, true)
	employee03 := employee.NewEmployee(3, "Juan", "Martínez", 80000, true)
	employee04 := employee.NewEmployee(4, "Laura", "Rodríguez", 100000, false)
	employee05 := employee.NewEmployee(5, "Carlos", "López", 110000, true)
	employee06 := employee.NewEmployee(6, "Ana", "Pérez", 95000, false)
	employee07 := employee.NewEmployee(7, "Pedro", "Sánchez", 85000, true)
	employee08 := employee.NewEmployee(8, "Sofía", "Hernández", 105000, true)
	employee09 := employee.NewEmployee(9, "Diego", "Gómez", 92000, false)
	employee10 := employee.NewEmployee(10, "Elena", "Díaz", 98000, true)

	service := employee.NewEmployeeService(employee.NewEmployeeRepository())

	service.Register(*employee01)
	service.Register(*employee02)
	service.Register(*employee03)
	service.Register(*employee04)
	service.Register(*employee05)
	service.Register(*employee06)
	service.Register(*employee07)
	service.Register(*employee08)
	service.Register(*employee09)
	service.Register(*employee10)

	for _, employee := range service.RetrieveAll(5) {
		fmt.Println(employee)
	}

}
