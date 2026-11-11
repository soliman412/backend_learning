package company

type PartTime struct {
	BaseEmployee
	hoursWorked int
	hourlyRate  float64
}

func NewPartTime(id int, name string, hoursWorked int, hourlyRate float64) PartTime {
	return PartTime{
		BaseEmployee: BaseEmployee{id: id, name: name},
		hoursWorked:  hoursWorked,
		hourlyRate:   hourlyRate,
	}
}

func (p PartTime) Salary() float64 {
	return float64(p.hoursWorked) * p.hourlyRate
}