package programs

import "time"

type Program struct {
	ID int
	Name string
	Description string
	DateTime time.Time
	UserID int
}

var programs = []Program{}

func (p Program) Save() {
programs = append(programs, p)
}