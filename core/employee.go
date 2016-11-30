package core

// Employee is a person hired by an organisation, assigned to one of the departments.
type Employee struct {
	ID             string
	FirstName      string
	LastName       string
	Level          int
	CurrentlyHired bool
}

// CanHire implements a business rule that tells that only employees of level 6 or above can hire.
func (e Employee) CanHire() bool {
	return e.Level > 5
}

// CanFire implements a business rule tht tells that only employees of level 6 or above can fire,
// but can fire only employees that are lower level than they are.
func (e Employee) CanFire(fired *Employee) bool {
	return e.Level > 5 && fired.Level < e.Level
}
