package employee

type Employee struct {
	id        int16
	firstname string
	lastname  string
	salary    float64
	active    bool
}

func NewEmployee(id int16, firstname string, lastname string, salary float64, active bool) *Employee {
	return &Employee{
		id:        id,
		firstname: firstname,
		lastname:  lastname,
		salary:    salary,
		active:    active,
	}
}

func (e Employee) Id() int16 {
	return e.id
}

func (e Employee) Firstname() string {
	return e.firstname
}

func (e Employee) Lastname() string {
	return e.lastname
}

func (e Employee) Salary() float64 {
	return e.salary
}

func (e Employee) IsActive() bool {
	return e.active
}

// Repository
type Employees interface {
	Persist(empoyee Employee) bool
	ExistsById(employeeId int16) bool
	GetById(employeeId int16) (*Employee, bool)
	List(limit int16) []Employee
}

// Use cases
type RegisterNewEmployeeUseCase interface {
	Register(employee Employee) error
}

type RetrieveEmployeeUseCase interface {
	RetrieveAll(limit int16) []Employee
	RetrieveById(employeeId int16) (*Employee, error)
}
