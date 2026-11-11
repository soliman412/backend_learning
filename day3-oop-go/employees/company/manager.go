package company

type Manager struct {
	FullTime
	bonus float64
}

func NewManager(id int, name string, monthlySalary float64, bonus float64) Manager {
	return Manager{
		FullTime: NewFullTime(id, name, monthlySalary),
		bonus:    bonus,
	}
}

func (m Manager) Salary() float64 {
	return m.FullTime.Salary() + m.bonus
}