package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/mysza/go-repo-reference/core"
)

// HRManager is an interface the web interface interacts with.
// Normally it will be replaced by the real-life HRManager struct (in production),
// but having an interface here allows for mocking in testing etc.
type HRManager interface {
	Hire(employeeID string, employee *core.Employee) (*core.Employee, error)
	Fire(managerID, employeeID string) (*core.Employee, error)
}

// HRManagerWebHandler is a struct holding reference to the use case implementation.
type HRManagerWebHandler struct {
	HRManager HRManager
}

// HireRequest is interfaces-only struct used to decode the web request.
type HireRequest struct {
	EmployingID       string
	EmployeeFirstName string
	EmployeeLastName  string
	EmployeeLevel     int
}

// Hire method implements the hire use case for a web interface.
func (handler HRManagerWebHandler) Hire(res http.ResponseWriter, req *http.Request) {
	var hireRequest HireRequest
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&hireRequest)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	employee := &core.Employee{
		FirstName: hireRequest.EmployeeFirstName,
		LastName:  hireRequest.EmployeeLastName,
		Level:     hireRequest.EmployeeLevel,
	}
	stored, err := handler.HRManager.Hire(hireRequest.EmployingID, employee)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	json.NewEncoder(res).Encode(stored)
}
