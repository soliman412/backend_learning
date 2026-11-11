package company

type Employee interface {
	ID() int
	Name() string
	Salary() float64
}

type BaseEmployee struct {
	id   int
	name string
}

func (b BaseEmployee) ID() int {
	return b.id
}

func (b BaseEmployee) Name() string {
	return b.name
}