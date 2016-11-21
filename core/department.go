package core

// Department is a part of organisation that employees are assigned to.
type Department struct {
	ID      string
	Name    string
	members []Employee
}

// type DepartmentRepository interface {
// 	Store(department Department) error
// 	FindByID(id string) (Department, error)
// }

// AddEmployee adds employee to the department.
// func (d *Department) AddEmployee(e Employee) error {
// 	if e.PayGrade > d.MinimumPayGrade {
// 		return fmt.Errorf("Cannot add employee with pay grade %v to a department with minimum pay grade %v", e.PayGrade, d.MinimumPayGrade)
// 	}
// 	d.members = append(d.members, e)
// 	return nil
// }

// // RemoveEmployee removes employee from the department.
// func (d *Department) RemoveEmployee(e Employee) {
// 	final := d.members[:0]
// 	for _, m := range d.members {
// 		if m.ID != e.ID {
// 			final = append(final, m)
// 		}
// 	}
// 	d.members = final
// }
