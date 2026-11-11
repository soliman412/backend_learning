package company

type FullTime struct {
	BaseEmployee
	monthlySalary float64
}

func NewFullTime(id int, name string, monthlySalary float64) FullTime {
	return FullTime{
		BaseEmployee:  BaseEmployee{id: id, name: name},
		monthlySalary: monthlySalary,
	}
}

func (f FullTime) Salary() float64 {
	return f.monthlySalary
}