package interfaces

import (
	"encoding/json"

	"github.com/mysza/go-repo-reference/core"
	"github.com/pborman/uuid"
)

type BoltEmployeeRepo struct {
	Database BoltDatabase
}

type BoltDatabase interface {
	GetEntity(entityKind string, key []byte) ([]byte, error)
	PutEntity(entityKind string, key, value []byte) error
}

type dbEmployee struct {
	FirstName      string
	LastName       string
	Level          int
	CurrentlyHired bool
}

func dbToCore(id string, e *dbEmployee) *core.Employee {
	return &core.Employee{
		ID:             id,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		Level:          e.Level,
		CurrentlyHired: e.CurrentlyHired,
	}
}

func coreToDB(e *core.Employee) (string, *dbEmployee) {
	id := e.ID
	empl := &dbEmployee{
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		Level:          e.Level,
		CurrentlyHired: e.CurrentlyHired,
	}
	// case when it's a new entity (no id)
	if len(id) == 0 {
		id = uuid.New()
	}
	return id, empl
}

func (r *BoltEmployeeRepo) FindByID(employeeID string) (*core.Employee, error) {
	// map to internal representation
	key := []byte(employeeID)
	value, err := r.Database.GetEntity("employee", key)
	if err != nil {
		return nil, err
	}
	var employee dbEmployee
	err = json.Unmarshal(value, &employee)
	if err != nil {
		return nil, err
	}
	return dbToCore(employeeID, &employee), nil
}

func (r *BoltEmployeeRepo) Store(employee *core.Employee) (*core.Employee, error) {
	id, e := coreToDB(employee)
	value, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	err = r.Database.PutEntity("employee", []byte(id), value)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
