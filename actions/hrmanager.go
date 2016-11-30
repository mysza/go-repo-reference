package actions

import (
	"fmt"

	"github.com/mysza/go-repo-reference/core"
)

type Logger interface {
	Info(toLog string)
	Warn(toLog string)
	Error(toLog string)
}

type EmployeeRepository interface {
	FindByID(employeeID string) (*core.Employee, error)
	Store(employee *core.Employee) (*core.Employee, error)
}

type HRManager struct {
	EmployeeRepo EmployeeRepository
	Log          Logger
}

func (hm *HRManager) Hire(employingID string, employee *core.Employee) (*core.Employee, error) {
	employing, err := hm.EmployeeRepo.FindByID(employingID)
	if err != nil {
		hm.Log.Error(fmt.Sprintf("Error getting employee by ID: %v", employingID))
		return nil, err
	}
	if !employing.CanHire() {
		return nil, fmt.Errorf("Employee %v cannot hire", employing.ID)
	}
	employee.CurrentlyHired = true
	return hm.EmployeeRepo.Store(employee)
}

func (hm *HRManager) Fire(managerID, employeeID string) (*core.Employee, error) {
	manager, err := hm.EmployeeRepo.FindByID(managerID)
	if err != nil {
		return nil, err
	}
	employee, err := hm.EmployeeRepo.FindByID(employeeID)
	if err != nil {
		return nil, err
	}
	if !manager.CanFire(employee) {
		return employee, fmt.Errorf("Employee %v cannot fire %v", managerID, employeeID)
	}
	employee.CurrentlyHired = false
	return hm.EmployeeRepo.Store(employee)
}
