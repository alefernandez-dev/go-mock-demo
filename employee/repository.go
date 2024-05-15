package employee

type EmployeeMapRepository struct {
	employees map[int16]Employee
}

func NewEmployeeRepository() *EmployeeMapRepository {
	return &EmployeeMapRepository{
		employees: make(map[int16]Employee),
	}
}

func (emr *EmployeeMapRepository) Persist(employee Employee) bool {
	emr.employees[employee.id] = employee
	return true
}

func (emr *EmployeeMapRepository) ExistsById(employeeId int16) bool {
	_, exists := emr.employees[employeeId]
	return exists
}

func (emr *EmployeeMapRepository) GetById(employeeId int16) (*Employee, bool) {
	employee, exists := emr.employees[employeeId]
	return &employee, exists
}

func (emr *EmployeeMapRepository) List(limit int16) []Employee {
	var employees []Employee
	count := int16(0)
	for _, employee := range emr.employees {
		if count >= limit {
			break
		}
		employees = append(employees, employee)
		count++
	}
	return employees
}
